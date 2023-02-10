package Organisation

import (
	"errors"
)

// AccountAttributes as provided in models.go
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

// Validate() checks to ensure that the required properties for API calls are defined and correct
func (aa AccountAttributes) Validate() error {
	// BankID
	if aa.BankID == "" {
		return errors.New("Account Attributes object missing BankID")
	}

	// BankIDCode
	if aa.BankIDCode == "" {
		return errors.New("Account Attributes object missing BankIDCode")
	}

	// Bic
	if aa.Bic == "" {
		return errors.New("Account Attributes object missing Bic")
	}

	// Country
	if aa.Country == nil {
		return errors.New("Account Attributes object missing Country")
	}

	// Name
	if len(aa.Name) == 0 {
		return errors.New("Account Attributes object missing Name")
	}
	if len(aa.Name) > 4 {
		return errors.New("Account Attributes object Name array length exceeds 4")
	}

	return nil
}
