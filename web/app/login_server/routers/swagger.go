//go:build swagger
// +build swagger

package routers

import (
	"github.com/gin-gonic/gin"
	//sg "github.com/njtc406/cityOfHope/web/app/login_server/docs/swagger"
	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	registerGroupHandler(defaultGroup, handler(initSwagger))
}

func initSwagger(r *gin.RouterGroup) {
	//sg.SwaggerInfo.Schemes = []string{`http`, `https`}
	//sg.SwaggerInfo.Version = `2.0`
	//sg.SwaggerInfo.Title = `登录服务`
	//sg.SwaggerInfo.BasePath = r.BasePath()
	//sg.SwaggerInfo.Description = sg.SwaggerInfo.Title + `-API`
	//sg.SwaggerInfo.InfoInstanceName = `login_server`
	r.GET("/swagger/*any", swaggerMiddle, ginSwagger.WrapHandler(swagFiles.Handler))
}

func swaggerMiddle(c *gin.Context) {
	sg.SwaggerInfo.Host = c.Request.Host
	c.Next()
}
