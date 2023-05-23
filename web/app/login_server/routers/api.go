// Package routers
// 模块名:
// 模块功能简介:
package routers

import (
	"github.com/gin-gonic/gin"
	ctl "github.com/njtc406/cityOfHope/web/app/login_server/controller"
)

func init() {
	registerGroupHandler(defaultGroup, initAPIs)
}

func initAPIs(r *gin.RouterGroup) {
	r.POST(`/login`, ctl.Login)       // 登录
	r.POST(`/create`, ctl.CreateUser) // 登录

	serverList := r.Group(`/server_list`)
	{
		auth := serverList.Use(ctl.UserToken)
		auth.GET(``, ctl.ServerList)
	}
	notice := r.Group(`/notice`)
	{
		auth := notice.Use(ctl.UserToken)
		auth.GET(``, ctl.Notice)
	}
}
