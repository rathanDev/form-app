package config

var version string = "/v1"

const accountUrl string = "/organisation/accounts"

var baseUrl = "http://interview-accountapi:8080"
// var baseUrl = "http://localhost:8080"

func SetBaseUrl(url string) {
	baseUrl = url
}

func AccountUrl() string {
	version = "/v1"
	var url = baseUrl + version + accountUrl
	return url
}
