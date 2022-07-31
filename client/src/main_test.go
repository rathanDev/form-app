package main

import (
	"form3-client/config"
	"form3-client/model"
	"form3-client/operation"
	"form3-client/util"
	"log"
	"testing"
)

const givenAccountNumber = "400300"
const givenAccountId = "eb0bd6f5-c3f5-44b2-b677-acd23cdde516"

var version int64 = 0
var versionPointer *int64 = &version

func TestInit(t *testing.T) {
	log.Println("----- ----- ----- Initiate API Testing ----- ----- -----")

	const baseUrl = "http://interview-accountapi:8080"
	// const baseUrl = "http://localhost:8080"
	config.SetBaseUrl(baseUrl)

	resp, err := operation.CheckHealth()
	if err != nil {
		t.Errorf("No error expected but %v", err)
	}

	const okayStatus = "200 OK"
	var actualStatus = resp.Status

	if actualStatus != okayStatus {
		t.Errorf("Operation status actual:%q expected:%q", actualStatus, okayStatus)
	}
}

func TestCreate_expect201Created(t *testing.T) {

	var accountData model.AccountData
	accountData.ID = givenAccountId
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = versionPointer

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = util.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = util.CreateBooleanPointer(false)
	accountAttr.AccountNumber = givenAccountNumber
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC"
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = util.CreateStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	resp, err := operation.Create(accountData)

	if err != nil {
		t.Errorf("No error expected but %v", err)
	}

	const createdStatus = "201 Created"
	var actualStatus = resp.Status

	if actualStatus != createdStatus {
		t.Errorf("Operation status actual:%q expected:%q", actualStatus, createdStatus)
	}
}

func TestCreate_expect409Conflict(t *testing.T) {

	var accountData model.AccountData
	accountData.ID = givenAccountId
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = versionPointer

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = util.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = util.CreateBooleanPointer(false)
	accountAttr.AccountNumber = givenAccountNumber
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC"
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = util.CreateStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	resp, err := operation.Create(accountData)

	if err != nil {
		t.Errorf("No error expected but %v", err)
	}

	const conflictStatus = "409 Conflict"
	var actualStatus = resp.Status

	if actualStatus != conflictStatus {
		t.Errorf("Operation status actual:%q expected:%q", actualStatus, conflictStatus)
	}
}

func TestCreate_expect400BadRequest_missingField(t *testing.T) {

	var accountData model.AccountData
	accountData.ID = givenAccountId
	// accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = versionPointer

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = util.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = util.CreateBooleanPointer(false)
	accountAttr.AccountNumber = givenAccountNumber
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC"
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = util.CreateStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	resp, err := operation.Create(accountData)

	if err != nil {
		t.Errorf("No error expected but %v", err)
	}

	const badRequestStatus = "400 Bad Request"
	var actualStatus = resp.Status

	if actualStatus != badRequestStatus {
		t.Errorf("Operation status actual:%q expected:%q", actualStatus, badRequestStatus)
	}
}

func TestCreate_expect400BadRequest_InvalidInput(t *testing.T) {

	var accountData model.AccountData
	accountData.ID = givenAccountId
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = versionPointer

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = util.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = util.CreateBooleanPointer(false)
	accountAttr.AccountNumber = givenAccountNumber
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC PQRS" // <--- incorrect format
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = util.CreateStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	resp, err := operation.Create(accountData)

	if err != nil {
		t.Errorf("No error expected but %v", err)
	}

	const badRequestStatus = "400 Bad Request"
	var actualStatus = resp.Status

	if actualStatus != badRequestStatus {
		t.Errorf("Operation status actual:%q expected:%q", actualStatus, badRequestStatus)
	}
}

func TestFetch(t *testing.T) {
	apiResponse := operation.Fetch()
	accountDataList := apiResponse.AccountDataList

	actualCount := len(accountDataList)
	const expectedCount int = 1

	if actualCount != expectedCount {
		t.Errorf("No of accounts actual:%q expected:%q", actualCount, expectedCount)
	}
}

func TestFetchMapped(t *testing.T) {
	accounts := operation.FetchMapped()

	if len(accounts) != 1 {
		t.Error("Expected one account")
	}

	account := accounts[0]
	var fetchedAccountNumber = account.AccountNumber

	if fetchedAccountNumber != givenAccountNumber {
		t.Errorf("AccountNumber expected:%q actual:%q", givenAccountNumber, fetchedAccountNumber)
	}
}

func TestDelete(t *testing.T) {
	resp, err := operation.Delete(givenAccountId, version)

	if err != nil {
		t.Errorf("Expected no err but actual:%v", err)
	}

	const noContentStatus = "204 No Content"
	var actualStatus = resp.Status

	if actualStatus != noContentStatus {
		t.Errorf("Operation status actual:%q expected:%q", actualStatus, noContentStatus)
	}
}
