package sap_api_wrapper

import (
	"encoding/json"
	"fmt"
)

type SapApiGetOrdersResult struct {
	Value []struct {
		DocNum   json.Number `json:"DocNum"`
		DocDate  string      `json:"DocDate"`
		OrderRef string      `json:"NumAtCard"`
	} `json:"value"`
	NextLink string `json:"odata.nextLink"`
}

type SapApiGetOrdersReturn struct {
	Body *SapApiGetOrdersResult
}

func SapApiGetOrders(params SapApiQueryParams) (SapApiGetOrdersReturn, error) {
	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticaed client")
		return SapApiGetOrdersReturn{}, err
	}

	resp, err := client.
		//DevMode().
		R().
		SetResult(SapApiGetOrdersResult{}).
		SetQueryParams(params.AsReqParams()).
		Get("Orders")
	if err != nil {
		fmt.Println(err)
		return SapApiGetOrdersReturn{}, err
	}

	return SapApiGetOrdersReturn{
		Body: resp.Result().(*SapApiGetOrdersResult),
	}, nil

}

func SapApiGetOrders_AllPages(params SapApiQueryParams) (SapApiGetOrdersReturn, error) {
	res := SapApiGetOrdersResult{}
	for page := 0; ; page++ {
		params.Skip = page * 20

		getItemsRes, err := SapApiGetOrders(params)
		if err != nil {
			return SapApiGetOrdersReturn{}, err
		}

		res.Value = append(res.Value, getItemsRes.Body.Value...)

		if getItemsRes.Body.NextLink == "" {
			break
		}
	}

	return SapApiGetOrdersReturn{
		Body: &res,
	}, nil
}
