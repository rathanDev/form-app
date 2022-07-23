package util

import (
	"net/http"
	"log"
)

func CreateBooleanPointer(val bool) *bool {
	return &val
}

func CreateStringPointer(val string) *string {
	return &val
}

func CreateNumberPointer(val int64) *int64 {
	return &val
}

func PrintHttpResponse(resp *http.Response) {
	log.Println("response:", resp)
	log.Println("response Status : ", resp.Status)
	log.Println("response Headers : ", resp.Header)
}
