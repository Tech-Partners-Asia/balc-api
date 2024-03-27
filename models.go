package balcapi

type (
	PayRequest struct {
		Amt         int    `json:"amt"`
		Description string `json:"description"`
	}
	Response struct {
		Data         interface{} `json:"data"`
		ErrorMessage string      `json:"errorMessage"`
	}
)
