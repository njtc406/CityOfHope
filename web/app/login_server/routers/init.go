// Package routers
// 模块名: 路由
// 功能描述: 用于设置web请求路由资源
// 作者:  yr  2023/4/22 0022 0:04
// 最后更新:  yr  2023/4/22 0022 0:04
package routers

import (
	"github.com/gin-gonic/gin"
)

type handler func(engine *gin.RouterGroup)

const (
	defaultGroup = `/api/v1`
)

var groupHandlerPool = make(map[string][]handler)

// RouteSet 路由注册
func RouteSet(e *gin.Engine) {
	//handlePool.Range(func(name, value interface{}) bool {
	//	api := e.Group(name.(string))
	//	if handle, ok := value.(handler); ok {
	//		handle(api)
	//	}
	//	return true
	//})

	for group, pool := range groupHandlerPool {
		api := e.Group(group)
		for _, hd := range pool {
			hd(api)
		}
	}
}

func registerGroupHandler(group string, hd handler) {
	_, ok := groupHandlerPool[group]
	if !ok {
		groupHandlerPool[group] = []handler{}
	}

	groupHandlerPool[group] = append(groupHandlerPool[group], hd)
}
