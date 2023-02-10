package Organisation

import (
	"testing"
)

// Constants for use in tests
const (
	ACCDATA_ID           string = "0d209d7f-d07a-4542-947f-5885fddddae2"
	ACCDATA_ORG_ID       string = "ba61483c-d5c5-4f50-ae81-6b8c039bea43"
	ACCDATA_ORG_TYPE     string = "accounts"
	ACCDATA_BANK_ID      string = "400300"
	ACCDATA_BANK_ID_CODE string = "GBDSC"
	ACCDATA_BIC          string = "NWBKGB22"
)

// TestAccDataCreateInstance() attempts to instantiate and validate an AccountData object
func TestAccDataCreateInstance(t *testing.T) {
	var version int64 = 0
	var country string = "GB"
	var name = []string{"paolo g"}
	account := AccountData{
		ID:             ACCDATA_ID,
		OrganisationID: ACCDATA_ORG_ID,
		Type:           ACCDATA_ORG_TYPE,
		Version:        &version,
		Attributes: &AccountAttributes{
			BankID:     ACCDATA_BANK_ID,
			BankIDCode: ACCDATA_BANK_ID_CODE,
			Bic:        ACCDATA_BIC,
			Country:    &country,
			Name:       name,
		},
	}

	// Try to Validate()
	if err := account.Validate(); err != nil {
		t.Fatal(err)
	}

	// Check values
	if account.ID != ACCDATA_ID {
		t.Errorf("Expected ID %s, received %s", ACCDATA_ID, account.ID)
	}
	if account.OrganisationID != ACCDATA_ORG_ID {
		t.Errorf("Expected OrganisationID %s, received %s", ACCDATA_ORG_ID, account.OrganisationID)
	}
	if account.Type != ACCDATA_ORG_TYPE {
		t.Errorf("Expected Type %s, received %s", ACCDATA_ORG_TYPE, account.Type)
	}
	if *account.Version != version {
		t.Errorf("Expected Version %d, received %d", version, *account.Version)
	}
}

// TestAccDataIncomplete() attempts to validate an incomplete AccountData object
func TestAccDataIncomplete(t *testing.T) {
	var version int64 = 0
	account := AccountData{
		ID:             ACCDATA_ID,
		OrganisationID: ACCDATA_ORG_ID,
		Type:           ACCDATA_ORG_TYPE,
		Version:        &version,
	}

	// Try to Validate()
	if err := account.Validate(); err == nil {
		t.Errorf("TestAccDataIncomplete() should have failed due to invalid AccountData{}")
	}
}
