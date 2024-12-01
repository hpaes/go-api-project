package main

import (
	"log"
	"net/http"

	"github.com/hpaes/go-api-project/src/api/controller"
	"github.com/hpaes/go-api-project/src/core/application/usecase"
	"github.com/hpaes/go-api-project/src/infrastructure/database"
	"github.com/hpaes/go-api-project/src/infrastructure/logger"
	"github.com/hpaes/go-api-project/src/infrastructure/repository"
)

func main() {
	dbConn, err := database.NewPqAdapter()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")

	accRepository := repository.NewAccountRepository(dbConn)
	log.Println("Account repository created")

	appLogger := &logger.ConsoleLogger{}
	getAccountUsecase := usecase.NewAccountUseCase(accRepository, appLogger)
	log.Println("Get account use case created")
	signupUsecase := usecase.NewSignupUseCase(accRepository, appLogger)
	log.Println("Signup use case created")

	signupController := controller.NewSignupController(signupUsecase, getAccountUsecase, appLogger)
	http.HandleFunc("/signup", signupController.Signup)
	http.HandleFunc("/account", signupController.GetAccount)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
