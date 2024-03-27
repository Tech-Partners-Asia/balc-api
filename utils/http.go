package utils

type (
	API struct {
		Url    string
		Method string
		Func   string
	}
)

const (
	TimeFormatYYYYMMDDHHMMSS = "20060102150405"
	TimeFormatYYYYMMDD       = "20060102"
	HttpContent              = "application/json"
	XmlContent               = "application/xml"
)
