package main

import (
	"github.com/glaydus/viesapi"
)

func main() {
	client := viesapi.NewVIESClient("", "")
	nip := "PL7272445205"

	status, err := client.GetAccountStatus()
	if status != nil {
		println(status.String())
	} else {
		println(err.Error())
	}

	data, err := client.GetVIESData(nip)
	if data != nil {
		println(data.String())
	} else {
		println(err.Error())
	}
}
