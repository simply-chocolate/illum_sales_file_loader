package sap_api_wrapper

import (
	"encoding/json"
	"fmt"
)

type SapApiGetInvoicesResult struct {
	Value []struct {
		DocNum   json.Number `json:"DocNum"`
		DocDate  string      `json:"DocDate"`
		OrderRef string      `json:"NumAtCard"`
	} `json:"value"`
	NextLink string `json:"odata.nextLink"`
}

type SapApiGetInvoicesReturn struct {
	Body *SapApiGetInvoicesResult
}

func SapApiGetInvoices(params SapApiQueryParams) (SapApiGetInvoicesReturn, error) {
	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticaed client")
		return SapApiGetInvoicesReturn{}, err
	}

	resp, err := client.
		//DevMode().
		R().
		SetResult(SapApiGetInvoicesResult{}).
		SetQueryParams(params.AsReqParams()).
		Get("Invoices")
	if err != nil {
		fmt.Println(err)
		return SapApiGetInvoicesReturn{}, err
	}

	return SapApiGetInvoicesReturn{
		Body: resp.Result().(*SapApiGetInvoicesResult),
	}, nil

}

func SapApiGetInvoices_AllPages(params SapApiQueryParams) (SapApiGetInvoicesReturn, error) {
	res := SapApiGetInvoicesResult{}
	for page := 0; ; page++ {
		params.Skip = page * 20

		getItemsRes, err := SapApiGetInvoices(params)
		if err != nil {
			return SapApiGetInvoicesReturn{}, err
		}

		res.Value = append(res.Value, getItemsRes.Body.Value...)

		if getItemsRes.Body.NextLink == "" {
			break
		}
	}

	return SapApiGetInvoicesReturn{
		Body: &res,
	}, nil
}
