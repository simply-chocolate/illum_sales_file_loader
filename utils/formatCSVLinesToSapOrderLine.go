package utils

import (
	"fmt"
	"illum_sales_file_loader/sap_api_wrapper"
	"strings"
	"time"
)

func formatCSVLinesToSapOrder(csvLines string) (string, sap_api_wrapper.SapApiOrderBody, error) {
	salesDataLines := strings.Split(csvLines, "\n")
	var sapOrderInstance sap_api_wrapper.SapApiOrderBody

	headerData := strings.Split(salesDataLines[0], ",")
	dateOfSale, err := time.Parse("20060102", headerData[0])
	if err != nil {
		return "", sap_api_wrapper.SapApiOrderBody{}, fmt.Errorf("couldn't parse the time for salesfile %v. error: ", headerData, err)
	}

	sapOrderInstance.DocDate = dateOfSale.Format("2006-01-02")
	sapOrderInstance.DocDueDate = dateOfSale.Format("2006-01-02")
	sapOrderInstance.CustomerCode = "100068"
	orderRef = headerData[0]+headerData[]+headerData[]+headerData[]+headerData[]+headerData[]
	sapOrderInstance.OrderRef = 
	//TODO:
	// Vi kan skrive dato+bonnr+tidspunkt+personalenr+kassenr
	// Hvis alle disse bliver lavet til en string og sat i order ref, og der så er en fil hvor den order ref vi laver nu passer med en der står i SAP, 
	// Så kan vi springe dem over.
	// Så vores invoice og order map vi skal tjekke exists på skal være disse. 

	for _, salesDataLine := range salesDataLines {
		salesData := strings.Split(salesDataLine, ",")

		sapOrderInstance.
			fmt.Println(salesDataLine)
	}
}
