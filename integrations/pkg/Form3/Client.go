package Form3

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"integrations/pkg/Form3/Organisation"
	"io/ioutil"
	"net/http"
)

// A Client object to store the Form3 API URL and an http.client instance to make calls to the API
type Client struct {
	Url        string       `json:"url,omitempty"`
	HttpClient *http.Client `json:"-"`
}

// Client object constructor
func New(Url string, HttpClient http.Client) (*Client, error) {
	client := &Client{
		Url: Url,
		HttpClient: &HttpClient,
	}

	return client, nil
}

// A request object for calls to the API with an account object
type AccountRequest struct {
	Data	Organisation.AccountData	`json:"data,omitempty"`
}

// A response object the Form3 API account response
type AccountResponse struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

// CreateAccount() attempts to call the API's account creation endpoint, in order to create an account using the
// provided AccountData object
func (c *Client) CreateAccount(account Organisation.AccountData) (Organisation.AccountData, error) {
	// Check that the input is valid
	if err := account.Validate(); err != nil {
		return Organisation.AccountData{}, err
	}

	// Build the request
	request := AccountRequest{
		Data: account,
	}
	request_json, err := json.Marshal(request)
	if err != nil {
		return Organisation.AccountData{}, err
	}
	request_buffer := bytes.NewBuffer([]byte(string(request_json)))

	// Request account creation
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/v1/organisation/accounts", c.Url), request_buffer)
	if err != nil {
		return Organisation.AccountData{}, err
	}
	req.Header.Add("Accept", `application/json`)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return Organisation.AccountData{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Organisation.AccountData{}, err
	}

	// Check if the request succeeded
	if resp.StatusCode != 201 {
		msg := fmt.Sprintf("Create Fail %d: %s, %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(body))
		return Organisation.AccountData{}, errors.New(msg)
	}

	// Extract the response data
	var body_obj AccountResponse
	if err := json.Unmarshal(body, &body_obj); err != nil {
		return Organisation.AccountData{}, err
	}

	data_json, err := json.Marshal(body_obj.Data)
	if err != nil {
		return Organisation.AccountData{}, err
	}

	// Extract the new account data
	var created_account Organisation.AccountData
	if err := json.Unmarshal(data_json, &created_account); err != nil {
		return Organisation.AccountData{}, err
	}

	return created_account, nil
}

// FetchAccount() attempts to fetch AccountData from the API using the provided account identifier string
func (c *Client) FetchAccount(id string) (Organisation.AccountData, error) {
	// Check that the input is valid
	if id == "" {
		return Organisation.AccountData{}, errors.New("Account ID required")
	}

	// Request account fetch using id
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/organisation/accounts/%s", c.Url, id), nil)
	if err != nil {
		return Organisation.AccountData{}, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return Organisation.AccountData{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Organisation.AccountData{}, err
	}

	// Check if the request succeeded
	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("Fetch Fail %d: %s, %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(body))
		return Organisation.AccountData{}, errors.New(msg)
	}

	// Extract the response data
	var body_obj AccountResponse
	if err := json.Unmarshal(body, &body_obj); err != nil {
		return Organisation.AccountData{}, err
	}

	data_json, err := json.Marshal(body_obj.Data)
	if err != nil {
		return Organisation.AccountData{}, err
	}

	// Extract the fetched account data
	var fetched_account Organisation.AccountData
	if err := json.Unmarshal(data_json, &fetched_account); err != nil {
		return Organisation.AccountData{}, err
	}

	return fetched_account, nil
}

// DeleteAccount() attempts to request deletion of an account via the API, using the provided AccountData object
func (c *Client) DeleteAccount(account Organisation.AccountData) error {
	// Check that the input is valid
	if err := account.Validate(); err != nil {
		return err
	}

	// Request account deletion using provided account object
	var url string = fmt.Sprintf("%s/v1/organisation/accounts/%s?version=%d", c.Url, account.ID, *account.Version)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check if the request succeeded
	if resp.StatusCode != 204 {
		msg := fmt.Sprintf("Delete Fail %d: %s, %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(body))
		return errors.New(msg)
	}

	return nil
}
