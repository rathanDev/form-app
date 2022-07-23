package model

type Account struct {
	ID             string
	OrganisationID string
	Type           string

	AccountNumber string
	BankID        string
	BankIDCode    string
	Country       string
	Name          []string
}

type ApiResponse struct {
	AccountDataList []AccountData `json:"data,omitempty"`
	Links           Links
}

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

type Links struct {
	First string
	Last  string
	Self  string
}

// {
// 	"data": [
// 	  {
// 		"attributes": {
// 		  "account_classification": "class1",
// 		  "account_matching_opt_out": true,
// 		  "account_number": "1O001",
// 		  "alternative_names": null,
// 		  "bank_id": "5O001",
// 		  "country": "Singapore",
// 		  "name": [
// 			"Jana"
// 		  ]
// 		},
// 		"created_on": "2022-01-03T00:00:00.000Z",
// 		"id": "f50fc252-03f3-11ed-a0bc-0242ac130101",
// 		"modified_on": "2022-02-03T00:00:00.000Z",
// 		"organisation_id": "f50fc252-03f3-11ed-a0bc-0242ac130201",
// 		"type": "accounts",
// 		"version": 1
// 	  }
// 	],
// 	"links": {
// 	  "first": "/v1/organisation/accounts?page%5Bnumber%5D=first",
// 	  "last": "/v1/organisation/accounts?page%5Bnumber%5D=last",
// 	  "self": "/v1/organisation/accounts"
// 	}
// }
