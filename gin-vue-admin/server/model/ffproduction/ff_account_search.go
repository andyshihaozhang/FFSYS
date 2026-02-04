package ffproduction

// FfAccountSearch 客户搜索结构体
type FfAccountSearch struct {
	AccountName string `json:"accountName" form:"accountName"`
	Phone       string `json:"phone" form:"phone"`
	Address     string `json:"address" form:"address"`
	PageInfo
}
