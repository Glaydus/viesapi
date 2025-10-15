package viesapi

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

type viesData struct {
	XMLName xml.Name  `xml:"result"`
	VIES    VIESData  `xml:"vies"`
	Error   ViesError `xml:"error"`
}

type viesAccountStatus struct {
	XMLName xml.Name    `xml:"result"`
	Account viesAccount `xml:"account"`
	Error   ViesError   `xml:"error"`
}

type viesAccount struct {
	UID         string          `xml:"uid"`
	Type        string          `xml:"type"`
	ValidTo     string          `xml:"validTo"`
	BillingPlan viesBillingPlan `xml:"billingPlan"`
	Requests    struct {
		VIESDataCount int `xml:"viesData"`
		TotalCount    int `xml:"total"`
	} `xml:"requests"`
}

type viesBillingPlan struct {
	Name              string  `xml:"name"`
	SubscriptionPrice float64 `xml:"subscriptionPrice"`
	ItemPrice         float64 `xml:"itemPrice"`
	ItemPriceStatus   float64 `xml:"itemPriceCheckStatus"`
	Limit             int     `xml:"limit"`
	RequestDelay      int     `xml:"requestDelay"`
	DomainLimit       int     `xml:"domainLimit"`
	OverPlanAllowed   bool    `xml:"overplanAllowed"`
	ExcelAddIn        bool    `xml:"excelAddin"`
	App               bool    `xml:"app"`
	CLI               bool    `xml:"cli"`
	Stats             bool    `xml:"stats"`
	Monitor           bool    `xml:"monitor"`
	FuncGetVIESData   bool    `xml:"funcGetVIESData"`
}

type ViesError struct {
	Code        int    `json:"code" xml:"code"`
	Description string `json:"description" xml:"description"`
}

type VIESClient struct {
	id      string
	key     string
	url     string
	errcode int
	errmsg  string
	err     Error
	uevat   EUVAT
	nip     NIP
}

const (
	numberEUVAT = iota
	numberNIP
)

const (
	production_url = "https://viesapi.eu/api"
	test_url       = "https://viesapi.eu/api-test"

	test_id  = "test_id"
	test_key = "test_key"
)

// Get VIES data for specified number
func (c *VIESClient) getData(euvat string) *VIESData {

	// clear error
	c.clear()

	// validate number and construct path
	suffix, ok := c.getPathSuffix(numberEUVAT, euvat)
	if !ok {
		return nil
	}

	//prepare url
	url := c.url + "/get/vies/" + suffix

	// send request
	res := c.get(url)
	if res == nil {
		c.set(CLI_CONNECT, "")
		return nil
	}

	// parse response
	var data viesData
	err := xml.Unmarshal(res, &data)
	if err != nil {
		c.set(CLI_RESPONSE, "")
		return nil
	}

	if data.Error.Code != 0 {
		c.set(data.Error.Code, data.Error.Description)
		return nil
	}

	return &data.VIES
}

// Get user account's status
func (c *VIESClient) getAccountStatus() *AccountStatus {

	// clear error
	c.clear()

	//prepare url
	url := c.url + "/check/account/status"

	// send request
	res := c.get(url)
	if res == nil {
		c.set(CLI_CONNECT, "")
		return nil
	}

	// parse response
	var data viesAccountStatus
	err := xml.Unmarshal(res, &data)
	if err != nil {
		c.set(CLI_RESPONSE, "")
		return nil
	}

	if data.Error.Code != 0 {
		c.set(data.Error.Code, data.Error.Description)
		return nil
	}

	return &AccountStatus{
		UID:               data.Account.UID,
		Type:              data.Account.Type,
		ValidTo:           c.getDateTime(data.Account.ValidTo),
		BillingPlanName:   data.Account.BillingPlan.Name,
		SubscriptionPrice: data.Account.BillingPlan.SubscriptionPrice,
		ItemPrice:         data.Account.BillingPlan.ItemPrice,
		ItemPriceStatus:   data.Account.BillingPlan.ItemPriceStatus,
		Limit:             data.Account.BillingPlan.Limit,
		RequestDelay:      data.Account.BillingPlan.RequestDelay,
		DomainLimit:       data.Account.BillingPlan.DomainLimit,
		OverPlanAllowed:   data.Account.BillingPlan.OverPlanAllowed,
		ExcelAddIn:        data.Account.BillingPlan.ExcelAddIn,
		App:               data.Account.BillingPlan.App,
		CLI:               data.Account.BillingPlan.CLI,
		Stats:             data.Account.BillingPlan.Stats,
		Monitor:           data.Account.BillingPlan.Monitor,
		FuncGetVIESData:   data.Account.BillingPlan.FuncGetVIESData,
		VIESDataCount:     data.Account.Requests.VIESDataCount,
		TotalCount:        data.Account.Requests.TotalCount,
	}
}

// Prepare authorization header content
func (c *VIESClient) auth(method, urlstr string) (string, bool) {

	//parse url
	url, _ := url.Parse(urlstr)
	if url == nil {
		c.set(CLI_INPUT, "")
		return "", false
	}
	host := url.Host
	port := "80"
	if url.Scheme == "https" {
		port = "443"
	}

	i := strings.LastIndexByte(host, ':')
	if i > 0 {
		host = host[:i]
		port = url.Host[i+1:]
	}

	// prepare auth header value
	nonce := c.randomHex(4)
	ts := time.Now().Unix()
	s := fmt.Sprintf("%d\n%s\n%s\n%s\n%s\n%s\n\n", ts, nonce, method, url.Path, host, port)

	mac := c.getMac(s)

	return fmt.Sprintf(`MAC id="%s", ts="%d", nonce="%s", mac="%s"`, c.id, ts, nonce, mac), true
}

// Prepare user agent information header content
func (c *VIESClient) userAgent() string {
	return fmt.Sprintf("VIESAPIClient/%s Go/%s", vies_version, runtime.GOOS)
}

// Clear error info
func (c *VIESClient) clear() {
	c.errcode = 0
	c.errmsg = ""
}

// Set error info
func (c *VIESClient) set(code int, msg string) {
	c.errcode = code
	if msg != "" {
		c.errmsg = msg
	} else {
		c.errmsg = c.err.message(code)
	}
}

// Get result of HTTP GET request
func (c *VIESClient) get(url string) []byte {

	auth, ok := c.auth("GET", url)
	if !ok {
		return nil
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	req.Header.Set("User-Agent", c.userAgent())
	req.Header.Set("Authorization", auth)

	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	return body
}

// Calculates HMAC256 from input string
func (c *VIESClient) getMac(input string) string {
	h := hmac.New(sha256.New, []byte(c.key))
	h.Write([]byte(input))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Get path suffix for specified number type
func (c *VIESClient) getPathSuffix(typ int, number string) (string, bool) {
	var b bool
	var path string

	if typ == numberNIP {
		if !c.nip.isValid(number) {
			c.set(CLI_NIP, "")
			return "", false
		}
		path, b = c.nip.normalize(number)
		path = "nip/" + path

	} else if typ == numberEUVAT {
		if !c.uevat.isValid(number) {
			c.set(CLI_EUVAT, "")
			return "", false
		}
		path, b = c.uevat.normalize(number)
		path = "euvat/" + path

	} else {
		c.set(CLI_NUMBER, "")
		return "", false
	}
	return path, b
}

func (c *VIESClient) getDateTime(str string) *time.Time {

	if str == "" {
		return nil
	}

	layout := "2006-01-02"

	if len(str) > 10 {
		layout += "T15:04:05"
	}
	if len(str) > 19 {
		switch c := str[19]; c {
		case '+', '-':
			layout += "-07:00"
		default:
			str = str[:19]
		}
	}

	t, err := time.Parse(layout, str)
	if err != nil {
		return nil
	}
	return &t
}

func (c *VIESClient) randomHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
