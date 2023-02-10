package Organisation

import (
	"testing"
)

// Constants for use in tests
const (
	ACCATTR_BANK_ID      string = "400300"
	ACCATTR_BANK_ID_CODE string = "GBDSC"
	ACCATTR_BIC          string = "NWBKGB22"
)

// TestAccAttrCreateInstance() attempts to instantiate and validate an AccountAttributes object
func TestAccAttrCreateInstance(t *testing.T) {
	var country string = "GB"
	var name = []string{"paolo g"}

	aa := AccountAttributes{
		BankID:     ACCATTR_BANK_ID,
		BankIDCode: ACCATTR_BANK_ID_CODE,
		Bic:        ACCATTR_BIC,
		Country:    &country,
		Name:       name,
	}

	// Try to Validate()
	if err := aa.Validate(); err != nil {
		t.Fatal(err)
	}

	// Check values
	if aa.BankID != ACCATTR_BANK_ID {
		t.Errorf("Expected BankID  %s, received %s", ACCATTR_BANK_ID, aa.BankID)
	}
	if aa.BankIDCode != ACCATTR_BANK_ID_CODE {
		t.Errorf("Expected BankID Code  %s, received %s", ACCATTR_BANK_ID_CODE, aa.BankIDCode)
	}
	if aa.Bic != ACCATTR_BIC {
		t.Errorf("Expected Bic %s, received %s", ACCATTR_BIC, aa.Bic)
	}
	if *aa.Country != country {
		t.Errorf("Expected Country %s, received %s", country, *aa.Country)
	}
}

// TestAccAttrIncomplete() attempts to validate an incomplete AccountAttributes object
func TestAccAttrIncomplete(t *testing.T) {
	aa := AccountAttributes{
		BankID:     ACCATTR_BANK_ID,
		BankIDCode: ACCATTR_BANK_ID_CODE,
		Bic:        ACCATTR_BIC,
	}

	// Try to Validate()
	if err := aa.Validate(); err == nil {
		t.Errorf("TestAccAttrIncomplete() should have failed due to invalid AccountAttributes{}")
	}
}
