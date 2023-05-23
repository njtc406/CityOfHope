// Package controller
// 模块名: 登录平台
// 模块功能简介: 接口流程控制
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/njtc406/chaosutil/validate"
	"github.com/njtc406/cityOfHope/web/app/login_server/define"
	"github.com/njtc406/cityOfHope/web/app/login_server/works"

	"net/http"
)

// Login 登录
//
//	@Description	登录
//	@Summary		登录
//	@Tags			登录
//	@Accept			json
//	@Param			{object}	body	define.ReqLogin	true	"请求参数"
//	@Produce		json
//	@Success		200	{object}	define.Response{Data=define.ResLoginInfo}	"响应参数"
//	@Router			/login [POST]
func Login(c *gin.Context) {
	res := define.NewResponse()
	p := new(define.ReqLogin)
	if err := validate.TransError(c.ShouldBindJSON(p)); err != nil {
		res.SetStatus(http.StatusBadRequest).SetMessage(err.Error())
		doResponse(c, res)
		return
	}

	works.Login(p, res)
	doResponse(c, res)
}

// CreateUser 创建账号
//
//	@Description	创建账号
//	@Summary		创建账号
//	@Tags			创建账号
//	@Accept			json
//	@Param			{object}	body	define.ReqCreateUser	true	"请求参数"
//	@Produce		json
//	@Success		200	{object}	define.Response{Data=define.ResLoginInfo}	"响应参数"
//	@Router			/create [POST]
func CreateUser(c *gin.Context) {
	res := define.NewResponse()
	p := new(define.ReqCreateUser)
	if err := validate.TransError(c.ShouldBindJSON(p)); err != nil {
		res.SetStatus(http.StatusBadRequest).SetMessage(err.Error())
		doResponse(c, res)
		return
	}

	works.CreateUser(p, res)
	doResponse(c, res)
}

// ServerList 服务器列表
//
//	@Description	服务器列表
//	@Summary		服务器列表
//	@Tags			服务器列表
//	@Produce		json
//	@Success		200	{object}	define.Response{Data=[]define.ServerInfo}	"响应参数"
//	@Router			/server_list [GET]
func ServerList(c *gin.Context) {
	res := define.NewResponse()
	p := new(define.ReqServerList)
	if err := validate.TransError(c.ShouldBindJSON(p)); err != nil {
		res.SetStatus(http.StatusBadRequest).SetMessage(err.Error())
		doResponse(c, res)
		return
	}

	works.ServerList(p, res)
	doResponse(c, res)
}

// Notice 系统公告
//
//	@Description	系统公告
//	@Summary		系统公告
//	@Tags			系统公告
//	@Accept			json
//	@Param			{object}	body	define.ReqNotice	true	"请求参数"
//	@Produce		json
//	@Success		200	{object}	define.Response{Data=[]define.NoticeInfo}	"响应参数"
//	@Router			/notice [GET]
func Notice(c *gin.Context) {
	res := define.NewResponse()
	p := new(define.ReqNotice)
	if err := validate.TransError(c.ShouldBindJSON(p)); err != nil {
		res.SetStatus(http.StatusBadRequest).SetMessage(err.Error())
		doResponse(c, res)
		return
	}

	works.Notice(p, res)
	doResponse(c, res)
}
