package operation

import (
	"form3-client/config"
	"net/http"
)

func CheckHealth() (*http.Response, error) {
	url := config.HealthUrl()
	resp, err := http.Get(url)
	return resp, err
}
