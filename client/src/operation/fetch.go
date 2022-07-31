package operation

import (
	"encoding/json"
	"fmt"
	"form3-client/config"
	"form3-client/model"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch() (model.ApiResponse, error) {
	url := config.AccountUrl()
	var apiResponse model.ApiResponse

	response, err := http.Get(url)
	if err != nil {
		return apiResponse, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return apiResponse, err
	}

	json.Unmarshal([]byte(string(responseData)), &apiResponse)
	return apiResponse, nil
}

func FetchMapped() ([]model.Account, error) {
	url := config.AccountUrl()
	var accounts []model.Account

	response, err := http.Get(url)
	if err != nil {
		return accounts, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return accounts, err
	}

	var apiResponse model.ApiResponse
	json.Unmarshal([]byte(string(responseData)), &apiResponse)
	accountDataList := apiResponse.AccountDataList
	accounts = mapAccounts(accountDataList)
	return accounts, nil
}

func mapAccounts(accountDataList []model.AccountData) []model.Account {
	var accounts []model.Account

	for _, accountData := range accountDataList {
		var account model.Account
		account.ID = accountData.ID
		account.OrganisationID = accountData.OrganisationID
		account.Type = accountData.Type

		account.AccountNumber = accountData.Attributes.AccountNumber
		account.BankID = accountData.Attributes.BankID
		account.BankIDCode = accountData.Attributes.BankIDCode
		account.Country = *accountData.Attributes.Country
		account.Name = accountData.Attributes.Name

		accounts = append(accounts, account)
	}

	return accounts
}

func printAccountDataList(accountDataList []model.AccountData) {
	fmt.Println("Print AccountData List")
	for i, val := range accountDataList {
		log.Println("i =>", i)

		log.Println("Attributes =>", *val.Attributes)
		log.Println("ID =>", val.ID)
		log.Println("OrganisationID =>", val.OrganisationID)
		log.Println("Type =>", val.Type)
		log.Println("Version =>", *val.Version)
	}
}
