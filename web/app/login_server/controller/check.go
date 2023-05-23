// Package controller
// 模块名: 模块名
// 功能描述: 描述
// 作者:  yr  2023/5/11 0011 21:24
// 最后更新:  yr  2023/5/11 0011 21:24
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/njtc406/cityOfHope/web/app/login_server/define"

	"net/http"
)

// UserToken token验证
func UserToken(c *gin.Context) {
	var res = define.NewResponse()
	checkInfo := &struct {
		Token     string
		UserID    string
		TimeStamp int64
	}{}
	err := c.ShouldBindHeader(checkInfo)
	if err != nil {
		res.SetStatus(http.StatusUnauthorized).SetMessage(`拒绝访问,请先登录`)
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	//ok := account.CheckToken(checkInfo.UserID, checkInfo.Token, checkInfo.TimeStamp)
	//if !ok {
	//	res.SetStatus(http.StatusUnauthorized).SetMessage(`拒绝访问,请先登录`)
	//	c.AbortWithStatusJSON(http.StatusUnauthorized, res)
	//	return
	//}

	c.Next()
}
