package handler

import (
	"net/http"

	"github.com/hpaes/go-api-project/src/core/errors"
	"github.com/hpaes/go-api-project/src/infrastructure/logger"
)

func HandleError(w http.ResponseWriter, err error, logger logger.LogHandler) {
	var status int
	switch err.(type) {
	case errors.InvalidHttpMethodErr:
		status = http.StatusMethodNotAllowed
		http.Error(w, "Invalid HTTP method", status)
	case errors.AccountNotFoundErr:
		status = http.StatusNotFound
		http.Error(w, "Account not found", status)
	case errors.InvalidRequestPayloadErr:
		status = http.StatusBadRequest
		http.Error(w, "Invalid request payload", status)
	case errors.InvalidQueryParamPayloadErr:
		status = http.StatusBadRequest
		http.Error(w, "Invalid query param", status)
	default:
		status = http.StatusInternalServerError
		http.Error(w, "An internal server error occured", status)
	}

	logger.LogError("Error occurred", err)
}
