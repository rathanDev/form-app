package config

const baseUrl string = "http://localhost:8080"
var version string = "/v1";
const accountUrl string = "/organisation/accounts"

func AccountUrl() string {
	version = "/v1"
	var url = baseUrl + version + accountUrl
	return url
}