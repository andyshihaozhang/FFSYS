package response

import "gin-vue-admin/server/model/ffproduction"

type FfAccountResponse struct {
	Account ffproduction.FfAccount `json:"account"`
}
