package operation

import (
	"form3-client/config"
	"form3-client/model"
	"form3-client/util"
	"log"
	"testing"
)

func TestCreateFetchDelete(t *testing.T) {

	const localhostUrl = "http://localhost:8080"
	config.SetBaseUrl(localhostUrl)

	const givenAccountNumber = "400300"
	const givenAccountId = "eb0bd6f5-c3f5-44b2-b677-acd23cdde516"
	var givenVersion = util.CreateNumberPointer(0)

	var accountData model.AccountData
	accountData.ID = givenAccountId
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = givenVersion

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

	Create(accountData)

	accounts := FetchMapped()
	log.Println("AccountsAfterCreation:", accounts)

	if len(accounts) != 1 {
		t.Error("Expected one account")
	}

	account := accounts[0]
	var fetchedAccountNumber = account.AccountNumber

	if fetchedAccountNumber != givenAccountNumber {
		t.Errorf("got %q, wanted %q", fetchedAccountNumber, givenAccountNumber)
	}

	Delete(givenAccountId, *givenVersion)

	accountsAfterDeletion := FetchMapped()
	log.Println("AccountsAfterDeletion:", accountsAfterDeletion)

	if len(accounts) == 0 {
		t.Error("Expected no accounts")
	}

}
