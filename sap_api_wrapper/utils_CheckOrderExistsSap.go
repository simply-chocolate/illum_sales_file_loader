package sap_api_wrapper

import "fmt"

func CheckOrderExistsSap(orderNumber string) (bool, error) {
	resp, err := SapApiGetOrders_AllPages(SapApiQueryParams{
		Select:  []string{"DocNum"},
		Filter:  fmt.Sprintf("NumAtCard eq '%v' and CardCode eq '100068'", orderNumber),
		OrderBy: []string{"DocNum desc"},
	})
	if err != nil {
		return true, err
	}

	if len(resp.Body.Value) != 0 {
		return true, nil
	}

	return false, nil
}
