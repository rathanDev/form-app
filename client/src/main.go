package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"form3-client/model"
	"form3-client/operation"
	"form3-client/util"
)

func main() {

	log.Println("#--- form3-client - Start as service ---#")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("\nExecute create fetch and delete")
		create()
		fetch()
		delete()
		fmt.Fprintf(w, "Successfully executed, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}

func create() {
	log.Println("\n\nCreate")

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
}

func fetch() {
	log.Println("\n\nFetch")
	apiResponse := operation.Fetch()
	log.Println(apiResponse)

	log.Println("\nFetchMapped")
	accounts := operation.FetchMapped()
	log.Println(accounts)
}

func delete() {
	log.Println("\n\nDelete")
	var id = "eb0bd6f5-c3f5-44b2-b677-acd23cdde516"
	var version = int64(0)
	operation.Delete(id, version)
}
