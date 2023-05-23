// Package controller
// 模块名: 控制器
// 功能描述: 用于处理请求信息
// 作者:  yr  2023/4/22 0022 0:04
// 最后更新:  yr  2023/4/22 0022 0:04
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/njtc406/cityOfHope/internal/pkg/log"
	"net/http"
	"time"
)

func doResponse(c *gin.Context, data interface{}, code ...int) {
	log.SysLogger.Warnf("login res:%#v", data)
	var status = http.StatusOK
	if len(code) > 0 {
		status = code[0]
	}

	c.JSON(status, data)
}

// setCookie 设置Cookie
func setCookie(c *gin.Context, Name, Value string, dur time.Duration) {
	c.SetCookie(Name, Value, int(dur/time.Second), `/`, ``, c.Request.URL.Scheme == `https`, true)
}

// getCookie 设置Cookie
func getCookie(C *gin.Context, Name string) (value string) {
	value, _ = C.Cookie(Name)
	return value
}

// delCookie 删除Cookie
func delCookie(c *gin.Context, Name string) {
	c.SetCookie(Name, ``, -1, ``, ``, c.Request.URL.Scheme == `https`, true)
}
