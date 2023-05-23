// Package node
// 模块名: 模块名
// 功能描述: 描述
// 作者:  yr  2023/5/15 0015 22:14
// 最后更新:  yr  2023/5/15 0015 22:14
package node

import (
	"github.com/njtc406/chaosutil/profiler"
	chaosTitle "github.com/njtc406/chaosutil/title"
	config "github.com/njtc406/cityOfHope/configs/login_server"
	"github.com/njtc406/cityOfHope/configs/public"
	"github.com/njtc406/cityOfHope/internal/pkg/log"
	"github.com/njtc406/cityOfHope/internal/pkg/pid"
	"github.com/njtc406/cityOfHope/internal/pkg/services"
	"github.com/njtc406/cityOfHope/internal/pkg/title"

	"os"
	"os/signal"
	"syscall"
	"time"
)

var exitCh = make(chan os.Signal, 0)

func init() {
	// TODO 初始化
	signal.Notify(exitCh, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
}

func Start() {
	title.EchoTitle()
	chaosTitle.EchoTitle()
	// TODO 整个登录服务器的启动流程
	// 1. 首先启动rpc服务
	// 2. 连接master服务器(这里需要一个tcp连接)
	// 3. 向master注册
	// 4. 根据master注册成功回复信息内容初始化配置
	// 5. 根据配置启动其他服务(对游戏客户端的所有api都为http提供)
	// 6. 登录模块(收到客户端发送来的账号和密码后通过平台验证后,通过jwt授权方式,返回客户端一个token,这里就可以在jwt信息中带入部分玩家信息)
	// 7. 服务器列表模块(收到客户端请求后,首先做鉴权,然后根据客户端版本选择对应的服务器列表返回)
	// 8. 服务器公告模块(收到客户端请求后,首先做鉴权,然后根据客户端版本选择对应的公告版本返回)
	// 9. 登录服上只记录玩家最后登录时间,不记录登出,需要记录常用ip,在新的ip登录时可以给绑定邮箱发个邮件或者给个提示

	log.SysLogger.Info("begin start server")
	initNode()
	run()

	// 记录pid
	pid.RecordPID(public.Conf.ProjectName, config.SystemConf.NodeType, config.SystemConf.NodeID)
	defer pid.DelPID(public.Conf.ProjectName, config.SystemConf.NodeType, config.SystemConf.NodeID)

	running := true
	pProfilerTicker := new(time.Ticker)
	if config.SystemConf.ProfilerInterval > 0 {
		pProfilerTicker = time.NewTicker(config.SystemConf.ProfilerInterval)
	}

	for running {
		select {
		case sig := <-exitCh:
			log.SysLogger.Infof("received the signal: %v", sig)
			running = false
		case <-pProfilerTicker.C:
			profiler.Report()
		}
	}
	log.SysLogger.Info("begin stop server...")
	stop()
	log.SysLogger.Info("server stopped, program exited...")
	log.Close()
	chaosTitle.EchoByeBye()
}

func initNode() {
	// 先初始化自己需要的东西

	// TODO 这里之后需要关注一下需不需要顺序
	for name, service := range services.Services {
		log.SysLogger.Infof("service %s : init", name)
		service.Init()
	}
}

func run() {
	// TODO 这里之后需要关注一下需不需要顺序
	for name, service := range services.Services {
		log.SysLogger.Infof("service %s : start", name)
		service.Start()
	}
}

func stop() {
	// TODO 这里之后需要关注一下需不需要顺序
	for name, service := range services.Services {
		log.SysLogger.Infof("service %s : stop", name)
		service.Stop()
	}
}
