// Package log
// 模块名: 模块名
// 功能描述: 描述
// 作者:  yr  2023/4/17 0017 23:11
// 最后更新:  yr  2023/4/17 0017 23:11
package log

import (
	chaoslog "github.com/njtc406/chaosutil/log"
	"github.com/njtc406/cityOfHope/configs/public"
)

var (
	logger    *chaoslog.DefaultLogger
	SysLogger *chaoslog.Logger
)

func init() {
	logger, err := chaoslog.NewDefaultLogger(
		public.Conf.GetSystemLoggerFileName(),
		public.Conf.SystemLogger.MaxAge,
		public.Conf.SystemLogger.RotationTime,
		public.Conf.SystemLogger.Level,
		public.Conf.SystemLogger.Caller,
		false,
		public.Conf.SystemLogger.Color,
		public.IsDebug(),
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
