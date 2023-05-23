// Package log
// 模块名: 模块名
// 功能描述: 描述
// 作者:  yr  2023/4/17 0017 23:11
// 最后更新:  yr  2023/4/17 0017 23:11
package log

import (
	chaoslog "github.com/njtc406/chaosutil/log"
	config "github.com/njtc406/cityOfHope/configs/login_server"
)

var (
	logger    *chaoslog.DefaultLogger
	SysLogger *chaoslog.Logger
)

func init() {
	logger, err := chaoslog.NewDefaultLogger(
		config.SystemConf.GetSystemLoggerFileName(),
		config.SystemConf.SystemLogger.MaxAge,
		config.SystemConf.SystemLogger.RotationTime,
		config.SystemConf.SystemLogger.Level,
		config.SystemConf.SystemLogger.Caller,
		false,
		config.SystemConf.SystemLogger.Color,
		config.IsDebug(),
	)

	if err != nil {
		panic(err)
	}

	SysLogger = logger.Logger
}

func Close() {
	SysLogger.Info("release log")
	logger.Close()
}
