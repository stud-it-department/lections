package structures

import "github.com/shopspring/decimal"

type JSONStruct struct {
	Integer          int `json:"integer"`
	DifficultInteger int `json:"difficult_integer"`

	D *decimal.Decimal `json:"d"`
	F *float64         `json:"f"`
}
