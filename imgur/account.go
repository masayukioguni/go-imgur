package imgur

import (
	"fmt"
	"net/http"
)

// AccountService Imgur API docs: https://api.imgur.com/models/account
type AccountService struct {
	client *Client
}

// Account Model
type Account struct {
	ID            int     `json:"id,omitempty"`
	Url           string  `json:"url,omitempty"`
	Bio           string  `json:"bio,omitempty"`
	Reputation    float32 `json:"reputation,omitempty"`
	Created       int     `json:"created,omitempty"`
	ProExpiration bool    `json:"pro_expiration,omitempty"`
}

// AccountResponse
type AccountResponse struct {
	Success bool    `json:"success,omitempty"`
	Status  int     `json:"status,omitempty"`
	Data    Account `json:"data,omitempty"`
}

// You can request the account information about any user.
func (s *AccountService) Account(username string) (*Account, *http.Response, error) {
	u := fmt.Sprintf("/3/account/%v", username)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AccountResponse)
	resp, err := s.client.Do(req, ar)
	if err != nil {
		return nil, resp, err
	}

	return &ar.Data, resp, err
}
