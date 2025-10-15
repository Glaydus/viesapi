package viesapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewVIESClient(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		key     string
		wantURL string
	}{
		{"with credentials", "my_id", "my_key", production_url},
		{"empty credentials", "", "", test_url},
		{"partial credentials", "id", "", test_url},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewVIESClient(tt.id, tt.key)
			if c == nil {
				t.Fatal("NewVIESClient returned nil")
			}
			if c.url != tt.wantURL {
				t.Errorf("url = %s, want %s", c.url, tt.wantURL)
			}
		})
	}
}

func TestVIESClientGetAccountStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xml := `<?xml version="1.0" encoding="UTF-8"?>
<result>
	<account>
		<uid>test-uid</uid>
		<type>premium</type>
		<billingPlan>
			<name>Premium</name>
		</billingPlan>
		<requests>
			<viesData>50</viesData>
			<total>100</total>
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

	status, err := c.GetAccountStatus()
	if err != nil {
		t.Fatalf("GetAccountStatus returned error: %v", err)
	}
	if status.UID != "test-uid" {
		t.Errorf("UID = %s, want test-uid", status.UID)
	}
}

func TestVIESClientGetAccountStatusError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xml := `<?xml version="1.0" encoding="UTF-8"?>
<result>
	<account></account>
	<error>
		<code>102</code>
		<description>Auth error</description>
	</error>
</result>`
		w.Write([]byte(xml))
	}))
	defer server.Close()

	c := NewVIESClient("test_id", "test_key")
	c.SetUrl(server.URL)

	status, err := c.GetAccountStatus()
	if status != nil {
		t.Error("GetAccountStatus should return nil on error")
	}
	if err == nil {
		t.Fatal("GetAccountStatus should return error")
	}
	if err.Code != 102 {
		t.Errorf("error code = %d, want 102", err.Code)
	}
}

func TestVIESClientGetVIESData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xml := `<?xml version="1.0" encoding="UTF-8"?>
<result>
	<vies>
		<uid>test-uid</uid>
		<countryCode>PL</countryCode>
		<vatNumber>1234567890</vatNumber>
		<valid>true</valid>
		<traderName>Test Company</traderName>
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

	data, err := c.GetVIESData("PL7272445205")
	if err != nil {
		t.Fatalf("GetVIESData returned error: %v", err)
	}
	if data.CountryCode != "PL" {
		t.Errorf("CountryCode = %s, want PL", data.CountryCode)
	}
	if !data.Valid {
		t.Error("Valid = false, want true")
	}
}

func TestVIESClientGetVIESDataError(t *testing.T) {
	c := NewVIESClient("", "")
	data, err := c.GetVIESData("invalid")
	if data != nil {
		t.Error("GetVIESData should return nil for invalid number")
	}
	if err == nil {
		t.Fatal("GetVIESData should return error")
	}
	if err.Code != CLI_EUVAT {
		t.Errorf("error code = %d, want %d", err.Code, CLI_EUVAT)
	}
}

func TestGetLastError(t *testing.T) {
	c := NewVIESClient("", "")
	c.set(CLI_NIP, "test error")
	code, msg := c.GetLastError()
	if code != CLI_NIP {
		t.Errorf("code = %d, want %d", code, CLI_NIP)
	}
	if msg != "test error" {
		t.Errorf("msg = %s, want test error", msg)
	}
}

func TestSetUrl(t *testing.T) {
	c := NewVIESClient("id", "key")
	customURL := "https://custom.url/api"
	c.SetUrl(customURL)
	if c.url != customURL {
		t.Errorf("url = %s, want %s", c.url, customURL)
	}
}

func TestViesErrorError(t *testing.T) {
	e := &ViesError{Code: 22, Description: "EU VAT ID is invalid"}
	errStr := e.Error()
	if errStr == "" {
		t.Error("Error() returned empty string")
	}
	var decoded ViesError
	if err := json.Unmarshal([]byte(errStr), &decoded); err != nil {
		t.Errorf("Error() returned invalid JSON: %v", err)
	}
	if decoded.Code != 22 {
		t.Errorf("decoded code = %d, want 22", decoded.Code)
	}
}

func TestAccountStatusString(t *testing.T) {
	validTo := time.Now()
	status := &AccountStatus{
		UID:             "test-uid",
		Type:            "premium",
		ValidTo:         &validTo,
		BillingPlanName: "Premium",
		Limit:           1000,
	}
	str := status.String()
	if str == "" {
		t.Error("String() returned empty string")
	}
	var decoded AccountStatus
	if err := json.Unmarshal([]byte(str), &decoded); err != nil {
		t.Errorf("String() returned invalid JSON: %v", err)
	}
	if decoded.UID != "test-uid" {
		t.Errorf("decoded UID = %s, want test-uid", decoded.UID)
	}
}

func TestVIESDataString(t *testing.T) {
	data := &VIESData{
		UID:         "test-uid",
		CountryCode: "PL",
		VATNumber:   "1234567890",
		Valid:       true,
		TraderName:  "Test Company",
	}
	str := data.String()
	if str == "" {
		t.Error("String() returned empty string")
	}
	var decoded VIESData
	if err := json.Unmarshal([]byte(str), &decoded); err != nil {
		t.Errorf("String() returned invalid JSON: %v", err)
	}
	if decoded.CountryCode != "PL" {
		t.Errorf("decoded CountryCode = %s, want PL", decoded.CountryCode)
	}
}

func TestAccountStatusNilValidTo(t *testing.T) {
	status := &AccountStatus{
		UID:     "test-uid",
		ValidTo: nil,
	}
	str := status.String()
	var decoded AccountStatus
	if err := json.Unmarshal([]byte(str), &decoded); err != nil {
		t.Errorf("String() with nil ValidTo returned invalid JSON: %v", err)
	}
}
