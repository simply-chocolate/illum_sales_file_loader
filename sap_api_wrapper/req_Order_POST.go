package sap_api_wrapper

import (
	"fmt"
	"time"
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

func SapApiPostOrder(OrderBody SapApiOrderBody, count int) error {
	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticated client")
		return err
	}
	resp, err := client.
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

	if resp.IsError() {
		orderExists, err := CheckOrderExistsSap(OrderBody.OrderRef)
		if err != nil {
			return fmt.Errorf("error checking if order exists after failing to post order to SAP. Order.nr: %v", OrderBody.OrderRef)
		}

		if orderExists {
			return nil
		} else if count > 200 {
			return fmt.Errorf("error posting order %v to sap. Tried %v times and still got nowhere", OrderBody.OrderRef, count)
		} else {
			time.Sleep(100 * time.Millisecond)
			SapApiPostOrder(OrderBody, count+1)
		}
	}

	return nil
}
