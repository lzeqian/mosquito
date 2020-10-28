package models

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
