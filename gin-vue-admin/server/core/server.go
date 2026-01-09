package core

import (
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
)

// ANSI 颜色代码
const (
	colorReset  = "\x1b[0m"
	colorBright = "\x1b[1m"
	colorCyan   = "\x1b[36m"
	colorGreen  = "\x1b[32m"
	colorBlue   = "\x1b[34m"
	colorYellow = "\x1b[33m"
)

func colorText(text, color string) string {
	return color + text + colorReset
}

func RunServer() {
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	// redis
	if global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		if global.GVA_CONFIG.System.UseMultipoint {
			initialize.RedisList()
		}
	}
	// mongo
	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// jwt
	if global.GVA_DB != nil {
		system.LoadAll()
	}
	// 路由
	Router := initialize.Routers()
	// 打印启动信息
	printServerLogo(address)
	// 启动服务
	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}

func printServerLogo(address string) {
	logo := `
  ███████╗███████╗      ███████╗███████╗██████╗ ██╗   ██╗███████╗██████╗
  ██╔════╝██╔════╝      ██╔════╝██╔════╝██╔══██╗██║   ██║██╔════╝██╔══██╗
  █████╗  █████╗  █████╗███████╗█████╗  ██████╔╝██║   ██║█████╗  ██████╔╝
  ██╔══╝  ██╔══╝  ╚════╝╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝██╔══╝  ██╔══██╗
  ██║     ██║           ███████║███████╗██║  ██║ ╚████╔╝ ███████╗██║  ██║
  ╚═╝     ╚═╝           ╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝
`
	fmt.Println(colorText(logo, colorCyan))
	fmt.Println(colorBright + "  FF-SERVER Management System" + colorReset)
	fmt.Println(colorText("  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", colorGreen))
	fmt.Printf("  %s    │ %s\n", colorText("版本", colorYellow), colorText(global.Version, colorGreen))
	fmt.Printf("  %s │ %s\n", colorText("后端地址", colorYellow), colorText(fmt.Sprintf("http://127.0.0.1%s", address), colorBlue))
	fmt.Printf("  %s │ %s\n", colorText("API文档", colorYellow), colorText(fmt.Sprintf("http://127.0.0.1%s/swagger/index.html", address), colorBlue))
	fmt.Printf("  %s │ %s\n", colorText("MCP SSE", colorYellow), colorText(fmt.Sprintf("http://127.0.0.1%s%s", address, global.GVA_CONFIG.MCP.SSEPath), colorBlue))
	fmt.Println(colorText("  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━", colorGreen))
	fmt.Println(colorText("  FF-SERVER 启动成功!", colorGreen))
	fmt.Println()
}
