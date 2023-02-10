package Form3

import (
	"fmt"
	"integrations/pkg/Form3/Organisation"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Global vars to hold mock webserver
var (
	mux    *http.ServeMux
	server *httptest.Server
)

// Prepare HttpClient and end-to-end testing constructs
func Build() (http.Client, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	HttpClient := http.Client{
		Timeout: time.Second * 30,
	}

	return HttpClient, func() {
		server.Close()
	}
}

// Use flat files in testdata folder to mock responses
func GetFixture(path string, t *testing.T) string {
	data, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

/* Base case tests start: */

// Constants for use in base case tests
const (
	BASE_ID           string = "0d209d7f-d07a-4542-947f-5885fddddae2"
	BASE_ORG_ID       string = "ba61483c-d5c5-4f50-ae81-6b8c039bea43"
	BASE_ORG_TYPE     string = "accounts"
	BASE_BANK_ID      string = "400300"
	BASE_BANK_ID_CODE string = "GBDSC"
	BASE_BIC          string = "NWBKGB22"
)

// TestCreateAccount() checks an AccountData instance against the expected output of a call to the API. The
// response is mocked from a fixture.
func TestCreateAccount(t *testing.T) {
	HttpClient, Teardown := Build()
	defer Teardown()

	client, err := New(server.URL, HttpClient)
	if err != nil {
		t.Fatal(err)
	}

	mux.HandleFunc("/v1/organisation/accounts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, GetFixture("base_account_response.json", t))
	})

	var version int64 = 0
	var country string = "GB"
	var name = []string{"paolo g"}
	var new_account = Organisation.AccountData{
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

	created_account, err := client.CreateAccount(new_account)
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
		t.Errorf("Expected Country %s, received %s", BASE_BANK_ID, created_account.Attributes.BankID)
	}
	if created_account.Attributes.Bic != BASE_BIC {
		t.Errorf("Expected Country %s, received %s", BASE_BIC, created_account.Attributes.Bic)
	}
}

// TestFetchAccount() checks the response object from a mocked fetch. The response is mocked from a fixture.
func TestFetchAccount(t *testing.T) {
	HttpClient, Teardown := Build()
	defer Teardown()

	client, err := New(server.URL, HttpClient)
	if err != nil {
		t.Fatal(err)
	}

	mux.HandleFunc("/v1/organisation/accounts/"+BASE_ID, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, GetFixture("base_account_response.json", t))
	})

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

// TestDeleteAccount() checks the expected output of a delete request.
func TestDeleteAccount(t *testing.T) {
	HttpClient, Teardown := Build()
	defer Teardown()

	client, err := New(server.URL, HttpClient)
	if err != nil {
		t.Fatal(err)
	}

	mux.HandleFunc("/v1/organisation/accounts/"+BASE_ID, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

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
