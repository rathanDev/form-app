package main

import (
	"form-app/model"
	"form-app/operation"
	"form-app/util"
	"log"
)

func main() {
	testCreate()
	testFetch()
	testDelete()
}

func testCreate() {
	log.Println("Test Create")
	var accountData model.AccountData
	accountData.ID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde516"
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = util.CreateNumberPointer(0)

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = util.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = util.CreateBooleanPointer(false)
	accountAttr.AccountNumber = "1006"
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC"
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = util.CreateStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	operation.Create(accountData)
	log.Println("\n\n")
}

func testFetch() {
	log.Println("Test Fetch")
	apiResponse := operation.Fetch()
	log.Println(apiResponse)

	log.Println("\n")
	log.Println("FetchMapped")
	accounts := operation.FetchMapped()
	log.Println(accounts)
	log.Println("\n\n")
}

func testDelete() {
	log.Println("Test Delete")
	var id = "eb0bd6f5-c3f5-44b2-b677-acd23cdde516"
	var version = int64(0)
	operation.Delete(id, version)
	log.Println("\n\n")
}
