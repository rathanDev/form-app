package operation

import (
	"fmt"
	"form3-client/config"
	"form3-client/util"
	"log"
	"net/http"
)

func Delete(id string, version int64) {

	client := &http.Client{}

	deleteUrl := createDeleteUrl(id, version)

	req, err := http.NewRequest("DELETE", deleteUrl, nil)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	util.PrintHttpResponse(resp)
}

func createDeleteUrl(id string, version int64) string {
	url := config.AccountUrl()
	const (
		urlTemplate = `%s/%s?version=%d`
	)
	var deleteUrl = fmt.Sprintf(urlTemplate, url, id, version)
	log.Println(deleteUrl)
	return deleteUrl
}
