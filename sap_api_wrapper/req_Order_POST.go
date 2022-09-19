package sap_api_wrapper

import (
	"fmt"
)

type SapApiOrderBody struct {
	CustomerCode string `json:"CardCode"` // 100068
	DocDate      string `json:"DocDate"`
	DocDueDate   string `json:"DocDueDate"`
	OrderRef     string `json:"NumAtCard"`
	ItemLines    []SapApiPostOrderDocumentLine
}

type SapApiPostOrderDocumentLine struct {
	ItemCode string  `json:"ItemCode"`
	BarCode  string  `json:"BarCode"`
	Quantity float64 `json:"Quantity"`
	// Price, Discount and Tax
	UnitPrice       float64 `json:"UnitPrice"`
	LineTotal       float64 `json:"LineTotal"`
	DiscountPercent float64 `json:"DiscountPercent"`
	VatGroup        string  `json:"VatGroup"` // S1

	UoMEntry int    `json:"UoMEntry"` // -1: Manuelt, 1: Pcs, 2: Case,
	UoMCode  string `json:"UoMCode"`  // "Manuelt", "Pcs", "Case"

	// Whs and Cost
	WarehouseCode   string // 11
	AccountCode     string // 12400
	CostingCode     string // 11
	COGSCostingCode string // 11
	COGSAccountCode string // 22200
}

func SapApiPostOrder(OrderBody SapApiOrderBody) error {
	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticated client")
		return err
	}
	_, err = client.
		//DevMode().
		R().
		SetResult(SapApiGetOrdersResult{}).
		SetBody(map[string]interface{}{
			"CardCode":      OrderBody.CustomerCode,
			"DocDate":       OrderBody.DocDate,
			"DocDueDate":    OrderBody.DocDueDate,
			"NumAtCard":     OrderBody.OrderRef,
			"DocumentLines": OrderBody.ItemLines,
		}).
		Post("Orders")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
