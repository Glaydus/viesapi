package viesapi

import "time"

const vies_version = "1.2.4"

type VIESData struct {
	UID               string `json:"uid" xml:"uid"`
	CountryCode       string `json:"country_code" xml:"countryCode"`
	VATNumber         string `json:"vat_number" xml:"vatNumber"`
	Valid             bool   `json:"valid" xml:"valid"`
	TraderName        string `json:"trader_name" xml:"traderName"`
	TraderCompanyType string `json:"trader_company_type" xml:"traderCompanyType"`
	TraderAddress     string `json:"trader_address" xml:"traderAddress"`
	ID                string `json:"id" xml:"id"`
	Date              string `json:"date" xml:"date"`
	Source            string `json:"source" xml:"source"`
}

type AccountStatus struct {
	UID               string    `json:"uid"`
	Type              string    `json:"type"`
	ValidTo           time.Time `json:"valid_to"`
	BillingPlanName   string    `json:"billing_plan_name"`
	SubscriptionPrice float64   `json:"subscription_price"`
	ItemPrice         float64   `json:"item_price"`
	ItemPriceStatus   float64   `json:"item_price_status"`
	Limit             int       `json:"limit"`
	RequestDelay      int       `json:"request_delay"`
	DomainLimit       int       `json:"domain_limit"`
	OverPlanAllowed   bool      `json:"over_plan_allowed"`
	ExcelAddIn        bool      `json:"excel_add_in"`
	App               bool      `json:"app"`
	CLI               bool      `json:"cli"`
	Stats             bool      `json:"stats"`
	Monitor           bool      `json:"monitor"`
	FuncGetVIESData   bool      `json:"func_get_vies_data"`
	VIESDataCount     int       `json:"vies_data_count"`
	TotalCount        int       `json:"total_count"`
}

// Create new VIESClient instance with specified id and key or use test credentials
func NewVIESClient(id, key string) *VIESClient {

	url := production_url
	if id == "" || key == "" {
		id = test_id
		key = test_key
		url = test_url
	}
	return &VIESClient{
		id:    id,
		key:   key,
		url:   url,
		err:   Error{},
		nip:   NIP{},
		uevat: EUVAT{},
	}
}

// Get current account status
// GetAccountStatus returns account status or nil in case of error
func (c *VIESClient) GetAccountStatus() *AccountStatus {
	return c.getAccountStatus()
}

// Get VIES data for specified number from EU VIES system
// GetVIESData returns VIES data or nil in case of error
func (c *VIESClient) GetVIESData(euvat string) *VIESData {
	return c.getData(euvat)
}

// Get last error message
func (c *VIESClient) GetLastError() (int, string) {
	return c.errcode, c.errmsg

}

// Set non default service URL
func (c *VIESClient) SetUrl(url string) {
	c.url = url
}
