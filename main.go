package main

import (
	"illum_sales_file_loader/utils"
	"log"

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

}
