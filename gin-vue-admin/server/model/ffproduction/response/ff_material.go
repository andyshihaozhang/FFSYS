package response

import "gin-vue-admin/server/model/ffproduction"

type FfMaterialResponse struct {
	Material ffproduction.FfMaterial `json:"material"`
}
