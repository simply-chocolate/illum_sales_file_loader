package main

import (
	"fmt"
	"illum_sales_file_loader/sap_api_wrapper"
	"illum_sales_file_loader/utils"
	"time"
)

func main() {
	utils.LoadEnv()

	fmt.Printf("%v: Started the Script \n", time.Now().UTC().Format("2006-01-02 15:04:05"))
	utils.SendUnknownErrorToTeams(fmt.Errorf("Started the Script"))
	err := utils.CreateOrdersSap()
	if err != nil {
		utils.SendUnknownErrorToTeams(err)
	}
	sap_api_wrapper.SapApiPostLogout()

	fmt.Printf("%v: Success \n", time.Now().UTC().Format("2006-01-02 15:04:05"))
}
