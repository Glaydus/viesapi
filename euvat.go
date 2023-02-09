package viesapi

import (
	"regexp"
	"strings"
)

// EU VAT number verificator
type EUVAT struct{}

// Normalizes form of the VAT number
func (e *EUVAT) normalize(number string) (string, bool) {

	if number == "" {
		return "", false
	}
	number = strings.NewReplacer("-", "", " ", "").Replace(number)
	if !regexp.MustCompile(`[A-Z]{2}[A-Z0-9+*]{2,12}`).MatchString(number) {
		return "", false
	}
	return number, true
}

// Checks if specified VAT number is valid
func (e *EUVAT) isValid(number string) bool {
	number, ok := e.normalize(number)
	if !ok {
		return false
	}
	cc := number[:2]
	num := number[2:]
	pattern, ok := cmap[cc]
	if !ok {
		return false
	}
	if !regexp.MustCompile(pattern).MatchString(number) {
		return false
	}

	if cc == "PL" {
		nip := NIP{}
		return nip.isValid(num)
	}
	return true
}

var cmap = map[string]string{
	"AT": "ATU\\d{8}",
	"BE": "BE[0-1]{1}\\d{9}",
	"BG": "BG\\d{9,10}",
	"CY": "CY\\d{8}[A-Z]{1}",
	"CZ": "CZ\\d{8,10}",
	"DE": "DE\\d{9}",
	"DK": "DK\\d{8}",
	"EE": "EE\\d{9}",
	"EL": "EL\\d{9}",
	"ES": "ES[A-Z0-9]{1}\\d{7}[A-Z0-9]{1}",
	"FI": "FI\\d{8}",
	"FR": "FR[A-Z0-9]{2}\\d{9}",
	"HR": "HR\\d{11}",
	"HU": "HU\\d{8}",
	"IE": "IE[A-Z0-9+*]{8,9}",
	"IT": "IT\\d{11}",
	"LT": "LT\\d{9,12}",
	"LU": "LU\\d{8}",
	"LV": "LV\\d{11}",
	"MT": "MT\\d{8}",
	"NL": "NL[A-Z0-9+*]{12}",
	"PL": "PL\\d{10}",
	"PT": "PT\\d{9}",
	"RO": "RO\\d{2,10}",
	"SE": "SE\\d{12}",
	"SI": "SI\\d{8}",
	"SK": "SK\\d{10}",
	"XI": "XI[A-Z0-9]{5,12}",
}
