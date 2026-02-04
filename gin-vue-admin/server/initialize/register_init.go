package initialize

import (
	_ "gin-vue-admin/server/source/example"
	_ "gin-vue-admin/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
