package operation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"form3-client/config"
	"form3-client/model"
	"net/http"
)

func Create(accountData model.AccountData) (*http.Response, error) {
	url := config.AccountUrl()
	payload := createPayload(accountData)
	resp, err := doPostToCreate(payload, url)
	return resp, err
}

func createPayload(accountData model.AccountData) []byte {
	a, _ := json.Marshal(accountData)
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

func doPostToCreate(payload []byte, url string) (*http.Response, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
