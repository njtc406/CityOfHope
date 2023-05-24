// Package login_server
// 模块名: web引擎
// 功能描述: 用于web模块的启动停止
// 作者:  yr  2023/4/22 0022 0:08
// 最后更新:  yr  2023/4/22 0022 0:08
package login_server

import (
	"context"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	config "github.com/njtc406/cityOfHope/configs/login_server"
	"github.com/njtc406/cityOfHope/configs/public"
	"github.com/njtc406/cityOfHope/internal/pkg/log"
	"github.com/njtc406/cityOfHope/internal/pkg/services"
	"github.com/njtc406/cityOfHope/web/app/login_server/routers"
	"github.com/njtc406/cityOfHope/web/app/login_server/works"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type serviceHttp struct {
	wg      *sync.WaitGroup
	running uint32
	handler *gin.Engine
	server  *http.Server
}

func (e *serviceHttp) Init() {
	// 替换验证器(这个东西之后再看用哪个版本)
	//*(binding.Validator.Engine().(*validator.Validate)) = *validate.Validator
	// 默认日志输出
	gin.DefaultWriter = io.MultiWriter(log.SysLogger.WriterLevel(logrus.InfoLevel))       // 设置默认日志输出为info级别
	gin.DefaultErrorWriter = io.MultiWriter(log.SysLogger.WriterLevel(logrus.ErrorLevel)) // 设置默认错误日志输出为error级别
	// 运行模式
	gin.SetMode(public.Conf.SystemStatus)
	// 默认中间件
	e.handler.Use(
		gzip.Gzip(gzip.DefaultCompression),
		gin.LoggerWithFormatter(e.logFormatter), // 这个设置的是默认日志的输出格式
		gin.Recovery(),
	)
	// 载入路由
	routers.RouteSet(e.handler)
	e.handler.ForwardedByClientIP = true
	// 初始化服务
	e.server = &http.Server{
		Addr:              config.SystemConf.HttpConf.Addr, // 服务监听端口
		Handler:           e.handler,
		ReadHeaderTimeout: time.Second * 3,
		IdleTimeout:       time.Second * 3,
	}
}

func (e *serviceHttp) logFormatter(p gin.LogFormatterParams) string {
	return fmt.Sprintf("[%s] %s %s %s %d %s \"%s\" %s\n",
		p.ClientIP,
		p.Method,
		p.Path,
		p.Request.Proto,
		p.StatusCode,
		p.Latency,
		p.Request.UserAgent(),
		p.ErrorMessage,
	)
}

func (e *serviceHttp) Start() {
	if !atomic.CompareAndSwapUint32(&e.running, 0, 1) {
		return
	}
	e.wg.Add(1)
	go e.run()
}

func (e *serviceHttp) run() {
	defer e.wg.Done()
	go works.Run()
	log.SysLogger.Infof("listen %s", e.server.Addr)
	if err := e.server.ListenAndServe(); err != nil {
		log.SysLogger.Error(err)
	}
}

func (e *serviceHttp) Stop() {
	defer atomic.StoreUint32(&e.running, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	works.Stop()

	if err := e.server.Shutdown(ctx); err != nil {
		log.SysLogger.Warn(err)
	}
	e.wg.Wait()
}

// NewService 创建新的HTTP服务
func NewService() *serviceHttp {
	return &serviceHttp{
		handler: gin.New(),
		server:  nil,
		wg:      new(sync.WaitGroup),
	}
}

func init() {
	services.Register(`login_service`, NewService())
}
