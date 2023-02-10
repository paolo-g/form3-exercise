package Organisation

import (
	"errors"
)

// AccountData as provided in models.go
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

// Validate() checks to ensure that the required properties for API calls are defined and correct
func (account AccountData) Validate() error {
	// Attributes
	if account.Attributes == nil {
		return errors.New("AccountData object missing Attributes")
	}

	err := account.Attributes.Validate()
	if err != nil {
		return err
	}

	// ID
	if account.ID == "" {
		return errors.New("AccountData object missing ID")
	}

	// OrganisationID
	if account.OrganisationID == "" {
		return errors.New("AccountData object missing Organisation ID")
	}

	// Type
	if account.Type == "" {
		return errors.New("AccountData object missing Type")
	}

	// Version
	if account.Version == nil {
		return errors.New("AccountData object missing Version")
	}

	return nil
}
