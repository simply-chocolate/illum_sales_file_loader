package utils

import (
	"fmt"
	"illum_sales_file_loader/sap_api_wrapper"
	"strconv"
	"strings"
	"time"
)

func formatCSVLinesAndPostOrder(csvLines string, ItemBarCodeCollection map[string]map[string]string) error {
	salesDataLines := strings.Split(csvLines, "\n")
	var sapOrderInstance sap_api_wrapper.SapApiOrderBody

	headerData := strings.Split(salesDataLines[0], ",")
	dateOfSale, err := time.Parse("20060102", headerData[0])
	if err != nil {
		return fmt.Errorf("couldn't parse the time for salesfile %v. error: %v ", headerData, err)
	}

	bookingDate := dateOfSale.Format("2006-01-02")
	sapOrderInstance.DocDate = bookingDate
	sapOrderInstance.DocDueDate = bookingDate
	sapOrderInstance.CustomerCode = "100068"
	orderRef := headerData[0] + headerData[4] + headerData[5] + headerData[6] + headerData[7]
	if len(orderRef) > 100 {
		orderRef = orderRef[:100]
	}

	sapOrders, err := GetOrdersFromSap(bookingDate)
	if err != nil {
		return fmt.Errorf("something went wrong getting the orders %v. error: %v", headerData, err)
	}
	if _, exists := sapOrders[orderRef]; exists {
		fmt.Println("order exists")
		return nil
	}

	sapInvoices, err := GetInvoicesFromSap(bookingDate)
	if err != nil {
		return fmt.Errorf("something went wrong getting the orders %v. error: %v", headerData, err)
	}
	if _, exists := sapInvoices[orderRef]; exists {
		fmt.Println("invoice exists")
		return nil
	}

	sapOrderInstance.OrderRef = orderRef

	for _, salesDataLine := range salesDataLines {
		salesData := strings.Split(salesDataLine, ",")
		//date := salesData[0]
		wareHouse := salesData[1]
		//brand := salesData[2]
		//timeOfDay := salesData[5]
		barCode := salesData[7]

		quantity, err := strconv.ParseFloat(salesData[8], 64)
		if err != nil {
			return fmt.Errorf("error parsing quantity as float. err: %v", err)
		}

		priceInclVat, err := strconv.ParseFloat(salesData[9], 64)
		if err != nil {
			return fmt.Errorf("error parsing price as float. err: %v", err)
		}

		discountInclVat, err := strconv.ParseFloat(strings.TrimSpace(salesData[10]), 64)
		if err != nil {
			fmt.Println(salesData[10])
			return fmt.Errorf("error parsing discount as float. err: %v", err)
		}

		itemBarCodeCollection, barCodeExists := ItemBarCodeCollection[barCode]
		if !barCodeExists {
			return fmt.Errorf("itemCode could not be found from barcode: %v", barCode)
		}

		uoMEntry, err := strconv.Atoi(itemBarCodeCollection["UoMEntry"])
		if err != nil {
			return fmt.Errorf("error converting UomEntry to int for barCode: %v err: %v ", barCode, err)
		}

		unitPrice := (priceInclVat * 0.8) / quantity / 100

		sapOrderInstance.ItemLines = append(sapOrderInstance.ItemLines, sap_api_wrapper.SapApiPostOrderDocumentLine{
			ItemCode:        itemBarCodeCollection["ItemCode"],
			UoMEntry:        uoMEntry,
			BarCode:         barCode,
			Quantity:        quantity,
			VatGroup:        "S1",
			UnitPrice:       unitPrice,
			DiscountPercent: discountInclVat / priceInclVat,
			LineTotal:       (priceInclVat * 0.8) / 100,

			WarehouseCode:   wareHouse,
			AccountCode:     "12400",
			CostingCode:     wareHouse,
			COGSAccountCode: FindCogsAccount(unitPrice),
			COGSCostingCode: wareHouse,
		})

	}

	err = sap_api_wrapper.SapApiPostOrder(sapOrderInstance)
	if err != nil {
		return fmt.Errorf("error posting order to SAP header: %v. Error: %v", headerData, err)
	}

	return nil
}
