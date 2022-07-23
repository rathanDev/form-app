package operation

import (
	"form3-client/config"
	"form3-client/model"
	"form3-client/util"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(accountData model.AccountData) {
	url := config.AccountUrl()

	payload := createPayload(accountData)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}

	util.PrintHttpResponse(resp)
}

func createPayload(accountData model.AccountData) []byte {
	a, _ := json.Marshal(accountData)
	log.Println(string(a))
	const (
		jsonTemplate = `{
			"data": %s
		  }
		`
	)
	var jsonPayload = fmt.Sprintf(jsonTemplate, a)
	var payload = []byte(jsonPayload)
	return payload
}
