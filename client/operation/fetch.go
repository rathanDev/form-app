package operation

import (
	"form-app/config"
	"form-app/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Fetch() model.ApiResponse {
	url := config.AccountUrl()
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var apiResponse model.ApiResponse
	json.Unmarshal([]byte(string(responseData)), &apiResponse)
	return apiResponse
}

func FetchMapped() []model.Account {
	url := config.AccountUrl()
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var apiResponse model.ApiResponse
	json.Unmarshal([]byte(string(responseData)), &apiResponse)
	accountDataList := apiResponse.AccountDataList
	var accounts []model.Account
	accounts = mapAccounts(accountDataList)
	return accounts
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
