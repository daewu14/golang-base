package test

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	grab2 "go_base_project/app/repositories/grab"
	"go_base_project/app/services/grab"
	"os"
	"testing"
)

func TestGetTokenService_Do(t *testing.T) {

	/*
		LOGRUS INIT
	*/
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("the err", err)
	}

	// call service
	service := grab.GetTokenService{
		Repo: grab2.GrabRepository{},
	}.Do()

	if service.Status != true {
		println("something went wrong ! ", service.Message)
		return
	}

	fmt.Println("success ", service.Data)
	return

}
