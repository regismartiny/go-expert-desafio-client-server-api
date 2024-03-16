package awesomeapi

import (
	"database/sql/driver"
	"encoding/json"
)

type Quotes struct {
	USDBRL Quote
}

type Quote struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func (t *Quotes) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), t)
}

func (t *Quotes) Value() (driver.Value, error) {
	b, err := json.Marshal(t)
	return string(b), err
}
