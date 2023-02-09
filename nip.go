package viesapi

import (
	"regexp"
	"strings"
)

// NIP number validator
type NIP struct{}

// Normalizes form of the NIP number
func (n *NIP) normalize(nip string) (string, bool) {
	if nip == "" {
		return "", false
	}
	nip = strings.NewReplacer("-", "", " ", "").Replace(nip)
	re := regexp.MustCompile(`[0-9]{10}`)
	if !re.MatchString(nip) {
		return "", false
	}
	return nip, true
}

// Checks if specified NIP is valid
func (n *NIP) isValid(nip string) bool {
	nip, ok := n.normalize(nip)
	if !ok {
		return false
	}

	w := [...]int{6, 5, 7, 2, 3, 4, 5, 6, 7}
	res := 0

	for i := range w {
		res += int(nip[i]-'0') * w[i]
	}
	res %= 11
	if res != int(nip[9]-'0') {
		return false
	}
	return true
}
