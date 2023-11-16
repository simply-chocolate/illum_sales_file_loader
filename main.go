package main

import (
	"fmt"
	"illum_sales_file_loader/sap_api_wrapper"
	"illum_sales_file_loader/utils"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	utils.LoadEnv()

	fmt.Printf("%v: Started the Script \n", time.Now().UTC().Format("2006-01-02 15:04:05"))
	err := utils.CreateOrdersSap()
	if err != nil {
		utils.SendUnknownErrorToTeams(err)
	}
	sap_api_wrapper.SapApiPostLogout()

	fmt.Printf("%v: Success \n", time.Now().UTC().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v: Started the Cron Scheduler", time.Now().UTC().Format("2006-01-02 15:04:05"))

	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Cron("0 * * * *").SingletonMode().Do(func() {
		fmt.Printf("%v: Started the Script\n", time.Now().Format("2006-01-02 15:04:05"))

		err := utils.CreateOrdersSap()
		if err != nil {
			utils.SendUnknownErrorToTeams(err)
		}

		sap_api_wrapper.SapApiPostLogout()
		fmt.Printf("%v: Success \n", time.Now().Format("2006-01-02 15:04:05"))
	})

	s.StartBlocking()
}
