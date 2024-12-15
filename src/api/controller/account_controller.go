package controller

import (
	"encoding/json"
	"net/http"

	"github.com/hpaes/go-api-project/src/api/handler"
	"github.com/hpaes/go-api-project/src/api/response"
	"github.com/hpaes/go-api-project/src/core/application/usecase"
	"github.com/hpaes/go-api-project/src/core/errors"
	"github.com/hpaes/go-api-project/src/infrastructure/logger"
)

type AccountController struct {
	su     usecase.SignUp
	ga     usecase.GetAccount
	logger logger.LogHandler
}

func NewSignupController(su usecase.SignUp, ga usecase.GetAccount, logger logger.LogHandler) *AccountController {
	return &AccountController{
		su:     su,
		ga:     ga,
		logger: logger,
	}
}

func (sc *AccountController) Signup(w http.ResponseWriter, r *http.Request) {
	sc.logger.LogInformation("AccountController.Signup")

	if r.Method != http.MethodPost {
		handler.HandleError(w, errors.NewInvalidHttpMethodErr(r.Method), sc.logger)
		return
	}

	var input usecase.SignupInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		handler.HandleError(w, errors.NewInvalidRequestPayloadErr("Input is not in a valid format"), sc.logger)
		return
	}

	signupOutput, err := sc.su.Execute(r.Context(), input)
	if err != nil {
		handler.HandleError(w, err, sc.logger)
		return
	}

	response.NewSuccessResponse(http.StatusCreated, signupOutput).Send(w)
}

func (sc *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {
	sc.logger.LogInformation("AccountController.GetAccount")
	if r.Method != http.MethodGet {
		handler.HandleError(w, errors.NewInvalidHttpMethodErr(r.Method), sc.logger)
		return
	}
	accountId := r.URL.Query().Get("account_id")
	if accountId == "" {
		handler.HandleError(w, errors.NewInvalidQueryParamErr("account_id"), sc.logger)
		return
	}

	getAccountOutput, err := sc.ga.Execute(r.Context(), accountId)
	if err != nil {
		handler.HandleError(w, errors.NewInternalServerErr("Error executing GetAccountUseCase", err), sc.logger)
		return
	}
	if getAccountOutput == nil {
		handler.HandleError(w, errors.NewAccountNotFoundErr(accountId), sc.logger)
		return
	}

	response.NewSuccessResponse(http.StatusOK, getAccountOutput).Send(w)
}
