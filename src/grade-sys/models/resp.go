package models

import(
	"errors"
)

type Data struct{
	Items []interface{} `json:"items"`
	RowCount int `json:"rowCount"`
}

type Resp struct {
	Data   `json:"data"`
	ErrorCode string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Code string `json:"code"`
	Message  string `json:"message"`
	Status string `json:"status"`
}

func (resp * Resp) Check() (int,error) {
	if resp.Code != "" && resp.Status == "error"{
		return -1,errors.New(resp.Message)
	}

	if resp.ErrorCode != "success" {
		return -1,errors.New(resp.ErrorMessage)
	}

	return resp.RowCount,nil
}