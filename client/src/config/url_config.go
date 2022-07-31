package config

var version string = "/v1"

const healthUrl string = "/health"
const accountUrl string = "/organisation/accounts"

var baseUrl = "http://interview-accountapi:8080"

// var baseUrl = "http://localhost:8080"

func SetBaseUrl(url string) {
	baseUrl = url
}

func HealthUrl() string {
	return baseUrl + version + healthUrl
}

func AccountUrl() string {
	return baseUrl + version + accountUrl
}
