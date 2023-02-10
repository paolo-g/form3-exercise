package main

import (
	"fmt"
	"integrations/pkg/Form3"
	"integrations/pkg/Form3/Organisation"
	"net/http"
	"os"
	"testing"
	"time"
)

// Prepare HttpClient
func Build() (string, http.Client) {
	Url := fmt.Sprintf("%s://%s:%s", os.Getenv("API_PROTOCOL"), os.Getenv("API_HOST"), os.Getenv("API_PORT"))
	HttpClient := http.Client{
		Timeout: time.Second * 30,
	}
	return Url, HttpClient
}

/* Base case tests start: */

// Constants for use in base case tests
const (
	BASE_ID           string = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	BASE_ORG_ID       string = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	BASE_ORG_TYPE     string = "accounts"
	BASE_BANK_ID      string = "123456"
	BASE_BANK_ID_CODE string = "GBDSC"
	BASE_BIC          string = "EXMPLGB2XXX"
)

// TestAccountCreateBase() instantiates an AccountData object using the BASE constants, calls the API via Docker and
// checks the result
func TestAccountCreateBase(t *testing.T) {
	Url, HttpClient := Build()

	client, err := Form3.New(Url, HttpClient)
	if err != nil {
		t.Fatal(err)
	}

	var version int64 = 0
	var country string = "GB"
	var name = []string{"paolo g"}
	var base_account = Organisation.AccountData{
		ID:             BASE_ID,
		OrganisationID: BASE_ORG_ID,
		Type:           BASE_ORG_TYPE,
		Version:        &version,
		Attributes: &Organisation.AccountAttributes{
			BankID:     BASE_BANK_ID,
			BankIDCode: BASE_BANK_ID_CODE,
			Bic:        BASE_BIC,
			Country:    &country,
			Name:       name,
		},
	}

	created_account, err := client.CreateAccount(base_account)
	if err != nil {
		t.Fatal(err)
	}

	// Check account properties
	if created_account.OrganisationID != BASE_ORG_ID {
		t.Errorf("Expected OrganisationID %s, received %s", BASE_ORG_ID, created_account.OrganisationID)
	}

	// Check account attributes properties
	if *created_account.Attributes.Country != country {
		t.Errorf("Expected Country %s, received %s", country, *created_account.Attributes.Country)
	}
	if created_account.Attributes.BankID != BASE_BANK_ID {
		t.Errorf("Expected Bank ID %s, received %s", BASE_BANK_ID, created_account.Attributes.BankID)
	}
	if created_account.Attributes.Bic != BASE_BIC {
		t.Errorf("Expected Bic %s, received %s", BASE_BIC, created_account.Attributes.Bic)
	}
}

// TestAccountFetchBase() calls the API (via Docker) using the BASE_ID constant, then checks the result
func TestAccountFetchBase(t *testing.T) {
	Url, HttpClient := Build()

	client, err := Form3.New(Url, HttpClient)
	if err != nil {
		t.Fatal(err)
	}

	fetched_account, err := client.FetchAccount(BASE_ID)
	if err != nil {
		t.Fatal(err)
	}

	// Check account properties
	if fetched_account.OrganisationID != BASE_ORG_ID {
		t.Errorf("Expected OrganisationID %s, received %s", BASE_ORG_ID, fetched_account.OrganisationID)
	}

	// Check account attributes properties
	if *fetched_account.Attributes.Country != "GB" {
		t.Errorf("Expected Country %s, received %s", "GB", *fetched_account.Attributes.Country)
	}
	if fetched_account.Attributes.BankID != BASE_BANK_ID {
		t.Errorf("Expected Bank ID %s, received %s", BASE_BANK_ID, fetched_account.Attributes.BankID)
	}
	if fetched_account.Attributes.Bic != BASE_BIC {
		t.Errorf("Expected Bic %s, received %s", BASE_BIC, fetched_account.Attributes.Bic)
	}
}

// TestAccountDeleteBase() instantiates an AccountData object using the BASE constants and calls the API via Docker
func TestAccountDeleteBase(t *testing.T) {
	Url, HttpClient := Build()

	client, err := Form3.New(Url, HttpClient)
	if err != nil {
		t.Fatal(err)
	}

	var version int64 = 0
	var country string = "GB"
	var name = []string{"paolo g"}
	var base_account = Organisation.AccountData{
		ID:             BASE_ID,
		OrganisationID: BASE_ORG_ID,
		Type:           BASE_ORG_TYPE,
		Version:        &version,
		Attributes: &Organisation.AccountAttributes{
			BankID:     BASE_BANK_ID,
			BankIDCode: BASE_BANK_ID_CODE,
			Bic:        BASE_BIC,
			Country:    &country,
			Name:       name,
		},
	}

	if err := client.DeleteAccount(base_account); err != nil {
		t.Fatal(err)
	}
}

/* End of base case tests */

/* Incomplete input tests start: */

// Constants for use in incomplete case tests
const (
	BAD_ID       string = "dfh93jf"
	BAD_ORG_ID   string = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	BAD_ORG_TYPE string = "accounts"
)

// TestAccountCreateIncomplete() attempts to instantiate an AccountData object with incomplete info
func TestAccountCreateIncomplete(t *testing.T) {
	Url, HttpClient := Build()

	client, err := Form3.New(Url, HttpClient)
	if err != nil {
		t.Fatal(err)
	}

	var bad_account = Organisation.AccountData{
		ID:             BAD_ID,
		OrganisationID: BAD_ORG_ID,
		Type:           BAD_ORG_TYPE,
	}

	if _, err := client.CreateAccount(bad_account); err == nil {
		t.Errorf("CreateAccount() should have failed due to missing input")
	}
}

// TestAccountFetchEmptyId() attempts to fetch using an empty string
func TestAccountFetchEmptyId(t *testing.T) {
	Url, HttpClient := Build()

	client, err := Form3.New(Url, HttpClient)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := client.FetchAccount(""); err == nil {
		t.Errorf("FetchAccount() should have failed due to missing ID")
	}
}

// TestAccountFetchFakeId() attempts to fetch using a nonexistant ID
func TestAccountFetchFakeId(t *testing.T) {
	Url, HttpClient := Build()

	client, err := Form3.New(Url, HttpClient)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := client.FetchAccount("h0h0h0"); err == nil {
		t.Errorf("FetchAccount() should have failed due to fake ID")
	}
}

// TestAccountDeleteIncomplete() attempts to delete an account using incomplete info
func TestAccountDeleteIncomplete(t *testing.T) {
	Url, HttpClient := Build()

	client, err := Form3.New(Url, HttpClient)
	if err != nil {
		t.Fatal(err)
	}

	var bad_account = Organisation.AccountData{
		OrganisationID: BAD_ORG_ID,
		Type:           BAD_ORG_TYPE,
	}

	if err := client.DeleteAccount(bad_account); err == nil {
		t.Errorf("DeleteAccount() should have failed due to missing input")
	}
}

/* End of incomplete input tests */
