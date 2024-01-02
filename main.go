package main

import (
	"fmt"
	"frozen-go-mini/route"
)

const (
	PORT = 7100
)

func main() {
	r := route.InitRouter()             // 注册路由
	_ = r.Run(fmt.Sprintf(":%d", PORT)) // 启动服务
}
