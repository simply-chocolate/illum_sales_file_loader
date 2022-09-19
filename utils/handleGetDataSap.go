package utils

import (
	"fmt"
	"illum_sales_file_loader/sap_api_wrapper"
	"strconv"
)

func GetInvoicesFromSap(date string) (map[string]string, error) {
	resp, err := sap_api_wrapper.SapApiGetInvoices_AllPages(sap_api_wrapper.SapApiQueryParams{
		Select:  []string{"DocNum", "DocDate", "NumAtCard"},
		Filter:  fmt.Sprintf("DocDate eq %v and CardCode eq '100068'", date),
		OrderBy: []string{"DocNum desc"},
	})
	if err != nil {
		return map[string]string{}, err
	}

	Invoices := make(map[string]string)

	for _, invoice := range resp.Body.Value {
		Invoices[invoice.BookingRef] = ""
	}

	return Invoices, nil
}

func GetOrdersFromSap(date string) (map[string]string, error) {
	resp, err := sap_api_wrapper.SapApiGetOrders_AllPages(sap_api_wrapper.SapApiQueryParams{
		Select:  []string{"DocNum", "DocDate", "NumAtCard"},
		Filter:  fmt.Sprintf("DocDate eq %v and CardCode eq '100068'", date),
		OrderBy: []string{"DocNum desc"},
	})
	if err != nil {
		return map[string]string{}, err
	}

	Orders := make(map[string]string)

	for _, order := range resp.Body.Value {
		Orders[order.BookingRef] = ""
	}

	return Orders, nil
}

func GetItemsFromSap() (map[string]map[string]string, error) {
	resp, err := sap_api_wrapper.SapApiGetItems_AllPages(sap_api_wrapper.SapApiQueryParams{
		Select: []string{"ItemCode", "ItemBarCodeCollection"},
		Filter: "Valid eq 'Y' and SalesItem eq 'Y'",
	})
	if err != nil {
		return map[string]map[string]string{}, err
	}

	BarCodes := make(map[string]map[string]string)

	BarCodes["POSTKORT"] = map[string]string{
		"ItemCode": "121",
		"UoMEntry": "-1",
	}

	for _, item := range resp.Body.Value {
		if len(item.ItemBarCodeCollection) == 0 {
			continue
		}
		for _, barCodeCollection := range item.ItemBarCodeCollection {
			if _, exists := BarCodes[barCodeCollection.Barcode]; !exists {
				BarCodes[barCodeCollection.Barcode] = map[string]string{
					"ItemCode": item.ItemCode,
					"UoMEntry": strconv.Itoa(barCodeCollection.UoMEntry),
				}
			}
		}
	}

	return BarCodes, nil
}
