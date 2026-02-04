package router

import (
	"gin-vue-admin/server/router/example"
	"gin-vue-admin/server/router/ffproduction"
	"gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	Ffproduction ffproduction.RouterGroup
}
