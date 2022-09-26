package main

import (
	"fmt"
	"illum_sales_file_loader/sap_api_wrapper"
	"illum_sales_file_loader/utils"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = utils.CreateOrdersSap()
	if err != nil {
		utils.SendUnknownErrorToTeams(err)
	}

	err = sap_api_wrapper.SapApiPostLogout()
	if err != nil {
		utils.SendUnknownErrorToTeams(err)
	}

	fmt.Printf("%v Success \n", time.Now().Format("2006-01-02 15:04:05"))
}
