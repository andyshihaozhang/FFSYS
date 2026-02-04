package response

import "gin-vue-admin/server/model/ffproduction"

type FfOperationResponse struct {
	Operation ffproduction.FfOperation `json:"operation"`
}
