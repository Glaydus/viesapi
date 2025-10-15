package viesapi

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetMac(t *testing.T) {
	c := NewVIESClient("test_id", "test_key")
	input := "test_input"
	mac := c.getMac(input)
	if mac == "" {
		t.Error("getMac returned empty string")
	}
	if mac != c.getMac(input) {
		t.Error("getMac not deterministic")
	}
}

func TestRandomHex(t *testing.T) {
	c := NewVIESClient("", "")
	hex := c.randomHex(4)
	if len(hex) != 8 {
		t.Errorf("randomHex(4) expected 8 chars, got %d", len(hex))
	}
}

func TestUserAgent(t *testing.T) {
	c := NewVIESClient("", "")
	ua := c.userAgent()
	if !strings.Contains(ua, "VIESAPIClient") {
		t.Error("userAgent missing VIESAPIClient")
	}
}

func TestClearAndSet(t *testing.T) {
	c := NewVIESClient("", "")
	c.set(CLI_NIP, "test error")
	if c.errcode != CLI_NIP || c.errmsg != "test error" {
		t.Error("set failed")
	}
	c.clear()
	if c.errcode != 0 || c.errmsg != "" {
		t.Error("clear failed")
	}
}

func TestAuth(t *testing.T) {
	c := NewVIESClient("test_id", "test_key")
	auth, ok := c.auth("GET", "https://viesapi.eu/api/test")
	if !ok {
		t.Error("auth failed")
	}
	if !strings.Contains(auth, "MAC id=") {
		t.Error("auth missing MAC id")
	}
	if !strings.Contains(auth, "test_id") {
		t.Error("auth missing client id")
	}
}

func TestGetPathSuffix(t *testing.T) {
	c := NewVIESClient("", "")

	tests := []struct {
		typ    int
		number string
		want   string
		ok     bool
	}{
		{numberEUVAT, "PL7272445205", "euvat/PL7272445205", true},
		{numberNIP, "7272445205", "nip/7272445205", true},
		{numberEUVAT, "invalid", "", false},
		{99, "1234567890", "", false},
	}

	for _, tt := range tests {
		got, ok := c.getPathSuffix(tt.typ, tt.number)
		if ok != tt.ok || (ok && got != tt.want) {
			t.Errorf("getPathSuffix(%d, %s) = %s, %v; want %s, %v", tt.typ, tt.number, got, ok, tt.want, tt.ok)
		}
	}
}

func TestGetDateTime(t *testing.T) {
	c := NewVIESClient("", "")

	tests := []struct {
		input string
		isNil bool
	}{
		{"", true},
		{"2024-01-15T10:30:45", false},
		{"2024-01-15T10:30:45+01:00", false},
		{"invalid", true},
	}

	for _, tt := range tests {
		result := c.getDateTime(tt.input)
		if (result == nil) != tt.isNil {
			t.Errorf("getDateTime(%s) nil=%v, want nil=%v", tt.input, result == nil, tt.isNil)
		}
	}
}

func TestGetData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xml := `<?xml version="1.0" encoding="UTF-8"?>
<result>
	<vies>
		<uid>test-uid</uid>
		<countryCode>PL</countryCode>
		<vatNumber>1234567890</vatNumber>
		<valid>true</valid>
	</vies>
	<error>
		<code>0</code>
		<description></description>
	</error>
</result>`
		w.Write([]byte(xml))
	}))
	defer server.Close()

	c := NewVIESClient("test_id", "test_key")
	c.SetUrl(server.URL)

	data := c.getData("PL7272445205")
	if data == nil {
		t.Error("getData returned nil")
	}
	if data != nil && data.CountryCode != "PL" {
		t.Errorf("getData countryCode = %s, want PL", data.CountryCode)
	}
}

func TestGetDataError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xml := `<?xml version="1.0" encoding="UTF-8"?>
<result>
	<vies></vies>
	<error>
		<code>22</code>
		<description>EU VAT ID is invalid</description>
	</error>
</result>`
		w.Write([]byte(xml))
	}))
	defer server.Close()

	c := NewVIESClient("test_id", "test_key")
	c.SetUrl(server.URL)

	data := c.getData("PL7272445205")
	if data != nil {
		t.Error("getData should return nil on error")
	}
	if c.errcode != 22 {
		t.Errorf("errcode = %d, want 22", c.errcode)
	}
}

func TestGetAccountStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xml := `<?xml version="1.0" encoding="UTF-8"?>
<result>
	<account>
		<uid>test-uid</uid>
		<type>premium</type>
		<validTo>2024-12-31T23:59:59</validTo>
		<billingPlan>
			<name>Premium</name>
			<subscriptionPrice>99.99</subscriptionPrice>
			<itemPrice>0.01</itemPrice>
			<itemPriceCheckStatus>0.02</itemPriceCheckStatus>
			<limit>10000</limit>
			<requestDelay>0</requestDelay>
			<domainLimit>5</domainLimit>
			<overplanAllowed>true</overplanAllowed>
			<excelAddin>true</excelAddin>
			<app>true</app>
			<cli>true</cli>
			<stats>true</stats>
			<monitor>true</monitor>
			<funcGetVIESData>true</funcGetVIESData>
		</billingPlan>
		<requests>
			<viesData>100</viesData>
			<total>150</total>
		</requests>
	</account>
	<error>
		<code>0</code>
		<description></description>
	</error>
</result>`
		w.Write([]byte(xml))
	}))
	defer server.Close()

	c := NewVIESClient("test_id", "test_key")
	c.SetUrl(server.URL)

	status := c.getAccountStatus()
	if status == nil {
		t.Fatal("getAccountStatus returned nil")
	}
	if status.UID != "test-uid" {
		t.Errorf("UID = %s, want test-uid", status.UID)
	}
	if status.Type != "premium" {
		t.Errorf("Type = %s, want premium", status.Type)
	}
	if status.VIESDataCount != 100 {
		t.Errorf("VIESDataCount = %d, want 100", status.VIESDataCount)
	}
}

func TestGetDataInvalidNumber(t *testing.T) {
	c := NewVIESClient("", "")
	data := c.getData("invalid")
	if data != nil {
		t.Error("getData should return nil for invalid number")
	}
	if c.errcode != CLI_EUVAT {
		t.Errorf("errcode = %d, want %d", c.errcode, CLI_EUVAT)
	}
}

func TestGetInvalidXML(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("invalid xml"))
	}))
	defer server.Close()

	c := NewVIESClient("test_id", "test_key")
	c.SetUrl(server.URL)

	data := c.getData("PL7272445205")
	if data != nil {
		t.Error("getData should return nil for invalid XML")
	}
	if c.errcode != CLI_RESPONSE {
		t.Errorf("errcode = %d, want %d", c.errcode, CLI_RESPONSE)
	}
}

func TestViesDataUnmarshal(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<result>
	<vies>
		<uid>test-uid</uid>
		<countryCode>PL</countryCode>
		<vatNumber>1234567890</vatNumber>
		<valid>true</valid>
	</vies>
	<error>
		<code>0</code>
		<description></description>
	</error>
</result>`

	var data viesData
	err := xml.Unmarshal([]byte(xmlData), &data)
	if err != nil {
		t.Fatalf("xml.Unmarshal failed: %v", err)
	}
	if data.VIES.CountryCode != "PL" {
		t.Errorf("CountryCode = %s, want PL", data.VIES.CountryCode)
	}
}

func TestViesAccountStatusUnmarshal(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<result>
	<account>
		<uid>test-uid</uid>
		<type>premium</type>
		<validTo>2024-12-31T23:59:59</validTo>
		<billingPlan>
			<name>Premium</name>
		</billingPlan>
		<requests>
			<viesData>100</viesData>
			<total>150</total>
		</requests>
	</account>
	<error>
		<code>0</code>
		<description></description>
	</error>
</result>`

	var data viesAccountStatus
	err := xml.Unmarshal([]byte(xmlData), &data)
	if err != nil {
		t.Fatalf("xml.Unmarshal failed: %v", err)
	}
	if data.Account.UID != "test-uid" {
		t.Errorf("UID = %s, want test-uid", data.Account.UID)
	}
}

func TestAuthWithPort(t *testing.T) {
	c := NewVIESClient("test_id", "test_key")
	auth, ok := c.auth("GET", "https://viesapi.eu:8443/api/test")
	if !ok {
		t.Error("auth failed with custom port")
	}
	if auth == "" {
		t.Error("auth returned empty string")
	}
}

func TestGetDateTimeWithTimezone(t *testing.T) {
	c := NewVIESClient("", "")
	result := c.getDateTime("2024-01-15T10:30:45+02:00")
	if result == nil {
		t.Error("getDateTime returned nil for valid datetime with timezone")
	}
	if result != nil {
		expected := time.Date(2024, 1, 15, 10, 30, 45, 0, time.FixedZone("", 2*3600))
		if !result.Equal(expected) {
			t.Errorf("getDateTime = %v, want %v", result, expected)
		}
	}
}
