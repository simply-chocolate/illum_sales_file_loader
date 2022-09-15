package sap_api_wrapper

import (
	"fmt"
)

type SapApiGetItemsResults struct {
	Value []struct {
		ItemCode              string `json:"ItemCode"`
		ItemBarCodeCollection []struct {
			Barcode  string `json:"Barcode"`
			UoMEntry int    `json:"UoMEntry"`
		} `json:"ItemBarCodeCollection"`
	} `json:"value"`
	NextLink string `json:"odata.nextLink"`
}

type SapApiGetItemsReturn struct {
	Body *SapApiGetItemsResults
}

func SapApiGetItems(params SapApiQueryParams) (SapApiGetItemsReturn, error) {
	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticaed client")
		return SapApiGetItemsReturn{}, err
	}

	resp, err := client.
		//DevMode().
		R().
		SetResult(SapApiGetItemsResults{}).
		SetQueryParams(params.AsReqParams()).
		Get("Items")
	if err != nil {
		fmt.Println(err)
		return SapApiGetItemsReturn{}, err
	}

	return SapApiGetItemsReturn{
		Body: resp.Result().(*SapApiGetItemsResults),
	}, nil

}

func SapApiGetItems_AllPages(params SapApiQueryParams) (SapApiGetItemsReturn, error) {
	res := SapApiGetItemsResults{}
	for page := 0; ; page++ {
		params.Skip = page * 20

		getItemsRes, err := SapApiGetItems(params)
		if err != nil {
			return SapApiGetItemsReturn{}, err
		}

		res.Value = append(res.Value, getItemsRes.Body.Value...)

		if getItemsRes.Body.NextLink == "" {
			break
		}
	}

	return SapApiGetItemsReturn{
		Body: &res,
	}, nil
}
