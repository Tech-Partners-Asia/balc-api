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

func New(endpoint, token string) Balc {
	return &balc{
		endpoint: endpoint,
		token:    token,
	}
}

func (b *balc) GetWebComponent(customerId int) string {
	return fmt.Sprintf("%s/?cust_id=%d&access_token=%s", b.endpoint, customerId, b.token)
}

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

func (b *balc) LimitCheck(customerId int) (LimitResponse, error) {
	var body []interface{} // empty array
	res, err := b.httpRequest(body, BalcLimit, customerId)
	if err != nil {
		return LimitResponse{}, err
	}
	var response LimitResponse
	json.Unmarshal(res, &response)

	return response, nil
}
