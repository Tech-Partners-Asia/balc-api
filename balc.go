package balcapi

import (
	"encoding/json"
)

type balc struct {
	endpoint string
	token    string
}

type Balc interface {
}

func New(endpoint, token string) Balc {
	return &balc{
		endpoint: endpoint,
		token:    token,
	}
}

func (b *balc) Loan(amount int, description string, customerId int) (Response, error) {
	var body []interface{}
	body = append(body, PayRequest{
		Amt:         amount,
		Description: description,
	})
	res, err := b.httpRequest(body, BalcLoan, customerId)
	if err != nil {
		return Response{}, err
	}
	var response Response
	json.Unmarshal(res, &response)

	return response, nil
}

func (b *balc) LimitCheck(customerId int) (Response, error) {
	var body []interface{} // empty array
	res, err := b.httpRequest(body, BalcLimit, customerId)
	if err != nil {
		return Response{}, err
	}
	var response Response
	json.Unmarshal(res, &response)

	return response, nil
}
