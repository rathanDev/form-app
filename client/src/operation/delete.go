package operation

import (
	"fmt"
	"form3-client/config"
	"net/http"
)

func Delete(id string, version int64) (*http.Response, error) {
	client := &http.Client{}

	deleteUrl := createDeleteUrl(id, version)

	req, err := http.NewRequest("DELETE", deleteUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	return resp, err
}

func createDeleteUrl(id string, version int64) string {
	url := config.AccountUrl()
	const (
		urlTemplate = `%s/%s?version=%d`
	)
	var deleteUrl = fmt.Sprintf(urlTemplate, url, id, version)
	return deleteUrl
}
