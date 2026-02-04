package response

import "gin-vue-admin/server/model/ffproduction"

type FfProductResponse struct {
	Product ffproduction.FfProduct `json:"product"`
}

type FfProductListResponse struct {
	ID              uint   `json:"ID"`
	ProductCode     string `json:"productCode"`
	ProductName     string `json:"productName"`
	ProductColor    string `json:"productColor"`
	ProductImg      string `json:"productImg"`
	AccountId       uint   `json:"accountId"`
	AccountName     string `json:"accountName"`
	ProductOrderNum int    `json:"productOrderNum"`
	ProductFlag     uint   `json:"productFlag"`
}
