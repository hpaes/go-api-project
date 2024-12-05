package handler

import (
	"net/http"

	"github.com/hpaes/go-api-project/src/core/errors"
	"github.com/hpaes/go-api-project/src/infrastructure/logger"
)

func HandleError(w http.ResponseWriter, err error, logger logger.LogHandler) {
	var status int
	var message string

	switch err.(type) {
	case errors.InvalidHttpMethodErr:
		status = http.StatusMethodNotAllowed
		message = "Invalid HTTP method"
	case errors.AccountNotFoundErr:
		status = http.StatusNotFound
		message = "Account not found"
	case errors.InvalidRequestPayloadErr:
		status = http.StatusBadRequest
		message = "Invalid request payload"
	case errors.InvalidQueryParamPayloadErr:
		status = http.StatusBadRequest
		message = "Invalid query param"
	case errors.AccountAlreadyExistsErr:
		status = http.StatusConflict
		message = "Account already exists"
	default:
		status = http.StatusInternalServerError
		message = "An internal server error occurred"
	}

	http.Error(w, message, status)
	logger.LogError("Error occurred: "+message, err)
}
