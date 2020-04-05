package error

import "net/http"

type AppErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(Message string) *AppErr {
	return &AppErr{Message: Message, Status: http.StatusBadRequest, Error: "bad_request"}
}

func NewNotFoundError(Message string) *AppErr {
	return &AppErr{Message: Message, Status: http.StatusBadRequest, Error: "not_found"}
}

func NewInternalServerError(Message string) *AppErr {
	return &AppErr{Message: Message, Status: http.StatusInternalServerError, Error: "internal_server_error"}
}
