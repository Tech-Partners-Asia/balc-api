package balcapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Tech-Partners-Asia/balc-api/utils"
)

var (
	BalcLoan = utils.API{
		Url:    "/api?custId=",
		Method: http.MethodPost,
		Func:   "loanadv",
	}
	BalcLimit = utils.API{
		Url:    "/api?custId=",
		Method: http.MethodPost,
		Func:   "limitcheck",
	}
)

func (b *balc) httpRequest(body interface{}, api utils.API, customerId int) (response []byte, err error) {
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	customerIdStr := strconv.Itoa(customerId)

	req, _ := http.NewRequest(api.Method, b.endpoint+api.Url+customerIdStr, requestBody)

	req.Header.Add("Content-Type", utils.HttpContent)
	req.Header.Add("Authorization", "Bearer "+b.token)
	req.Header.Add("func", api.Func)

	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		return nil, errors.New(string(res.Status))
	}
	defer res.Body.Close()
	response, _ = io.ReadAll(res.Body)
	return
}
