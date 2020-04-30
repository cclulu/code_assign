package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e HttpError) Error() string {
	return fmt.Sprintf("status_code:%d\nerror_message:%s\n", e.StatusCode, e.Message)
}

func StatusError(resp *http.Response) error {
	resp_err := HttpError{StatusCode: resp.StatusCode, Message: "empty response body"}
	data, err := ioutil.ReadAll(resp.Body)
	if err == nil && data != nil {
		resp_err.Message = string(data)
	}
	return resp_err
}
