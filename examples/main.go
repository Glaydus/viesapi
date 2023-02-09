package main

import (
	"encoding/json"

	"github.com/glaydus/viesapi"
)

func main() {
	client := viesapi.NewVIESClient("", "")
	nip := "PL7171642051"

	status := client.GetAccountStatus()
	if status != nil {

		b, _ := json.MarshalIndent(status, "", "    ")
		if b != nil {
			println(string(b))
		}
	}

	data := client.GetVIESData(nip)
	if data != nil {
		b, _ := json.MarshalIndent(data, "", "    ")
		if b != nil {
			println(string(b))
		}
	}

}
