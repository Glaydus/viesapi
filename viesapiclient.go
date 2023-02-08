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

type VIESData struct {
	UID               string `json:"uid"`
	CountryCode       string `json:"country_code"`
	VATNumber         string `json:"vat_number"`
	Valid             bool   `json:"valid"`
	TraderName        string `json:"trader_name"`
	TraderCompanyType string `json:"trader_company_type"`
	TraderAddress     string `json:"trader_address"`
	ID                string `json:"id"`
	Date              string `json:"date"`
	Source            string `json:"source"`
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
	vies_version = "1.2.5"

	numberEUVAT = iota
	numberNIP
)

func NewVIESClient(id, key string) *VIESClient {

	const (
		production_url = "https://viesapi.eu/api"
		test_url       = "https://viesapi.eu/api-test"

		test_id  = "test_id"
		test_key = "test_key"
	)

	Client := &VIESClient{}
	if id == "" || key == "" {
		Client.id = test_id
		Client.key = test_key
		Client.url = test_url
	} else {
		Client.id = id
		Client.key = key
		Client.url = production_url
	}
	Client.err = Error{}
	Client.uevat = EUVAT{}
	Client.nip = NIP{}

	return Client
}

// Get VIES data for specified number
func (c *VIESClient) GetVIESData(euvat string) (*VIESData, bool) {

	// clear error
	c.clear()

	// validate number and construct path
	suffix, ok := c.getPathSuffix(numberEUVAT, euvat)
	if !ok {
		return nil, false
	}

	//prepare url
	url := c.url + "/get/vies/" + suffix

	// send request
	res := c.get(url)
	if res == nil {
		c.set(CLI_CONNECT, "")
		return nil, false
	}

	var out any
	err := xml.Unmarshal(res, &out)
	if err != nil {
		c.set(CLI_RESPONSE, "")
		return nil, false
	}

	data := &VIESData{}

	return data, true
}

// Set non default service URL
func (c *VIESClient) SetUrl(url string) {
	c.url = url
}

// Get last error message
func (c *VIESClient) GetLastError() (int, string) {
	return c.errcode, c.errmsg
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
		port = host[i+1:]
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
		if !c.nip.IsValid(number) {
			c.set(CLI_NIP, "")
			return "", false
		}
		path, b = c.nip.normalize(number)

	} else if typ == numberEUVAT {
		if !c.uevat.IsValid(number) {
			c.set(CLI_EUVAT, "")
			return "", false
		}
		path, b = c.uevat.normalize(number)

	} else {
		c.set(CLI_NUMBER, "")
		return "", false
	}
	return path, b
}

func (c *VIESClient) randomHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
