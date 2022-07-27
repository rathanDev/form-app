package operation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"form3-client/config"
	"form3-client/model"
	// "form3-client/util"
	"log"
	"net/http"
)

func Create(accountData model.AccountData) *http.Response {
	url := config.AccountUrl()
	payload := createPayload(accountData)
	resp := doPostToCreate(payload, url)
	return resp
	// util.PrintHttpResponse(resp)
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

func doPostToCreate(payload []byte, url string) *http.Response {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
