package v1

import (
	"errors"
	"net/http"
)

func parseCtrlErr(err error) error {
	switch err.Error() {
	default:
		return errors.New("something went wrong")
	}
}

type WebErr struct {
	Status  int
	Message string
}

func (h WebErr) Error() string {
	return h.Message
}

func NewBadRequest(err error) WebErr {
	return WebErr{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}
