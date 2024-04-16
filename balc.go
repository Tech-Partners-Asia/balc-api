package balcapi

import (
	"encoding/json"
	"fmt"
)

type balc struct {
	endpoint string
	token    string
}

type Balc interface {
	Loan(amount int, description string, customerId int) (string, error)
	LimitCheck(customerId int) (LimitResponse, error)
	GetWebComponent(customerId int) string
}

// New creates a new Balc instance with the provided endpoint and token.
//
// Parameters:
// - endpoint: a string representing the endpoint
// - token: a string representing the token
//
// Return type: Balc
func New(endpoint, token string) Balc {
	return &balc{
		endpoint: endpoint,
		token:    token,
	}
}

// GetWebComponent generates a string representing a web component URL with the given customer ID.
//
// customerId: an integer representing the customer ID.
// Returns a string representing the web component URL.
func (b *balc) GetWebComponent(customerId int) string {
	return fmt.Sprintf("%s/?cust_id=%d&access_token=%s", b.endpoint, customerId, b.token)
}

// Loan is a function that performs a loan transaction.
//
// It takes the following parameters:
// - amount: an integer representing the loan amount
// - description: a string describing the loan
// - customerId: an integer representing the customer ID
//
// It returns a balc customer account ID and an error.
func (b *balc) Loan(amount int, description string, customerId int) (string, error) {
	var body []PayRequest
	body = append(body, PayRequest{
		Amt:         amount,
		Description: description,
	})
	res, err := b.httpRequest(body, BalcLoan, customerId)
	if err != nil {
		return "", err
	}
	var response string
	json.Unmarshal(res, &response)

	return response, nil
}

// LimitCheck checks the credit limit for a customer.
//
// It takes the customer ID as a parameter and returns a LimitResponse struct and an error.
//
// It returns a customer credit limit and an error.
func (b *balc) LimitCheck(customerId int) (LimitResponse, error) {
	body := make([]interface{}, 0)
	res, err := b.httpRequest(body, BalcLimit, customerId)
	if err != nil {
		return LimitResponse{}, err
	}
	var response LimitResponse
	json.Unmarshal(res, &response)

	return response, nil
}
