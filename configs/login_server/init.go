package config

import (
	"github.com/njtc406/chaosutil/validate"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

var (
	SystemConf = new(systemConf)
	parser     = viper.New()
)

func init() {
	// TODO 这里应该是需要修改的,一部分系统配置可能是由集群控制节点下发的

	setDefaultConf()
	parseSystemConfig()
	initDir()
}

func initDir() {
	if err := os.MkdirAll(SystemConf.LogPath, 0644); err != nil {
		panic(err)
	}
}

func setDefaultConf() {
	// 全局配置
	parser.SetDefault(`SystemStatus`, Debug)
	parser.SetDefault(`LogPath`, `./run/logs`)
	// 日志默认配置
	parser.SetDefault(`SystemLogger.Name`, `system.log`)
	parser.SetDefault(`SystemLogger.Level`, 5)
	parser.SetDefault(`SystemLogger.Caller`, false)
	parser.SetDefault(`SystemLogger.Color`, false)
	parser.SetDefault(`SystemLogger.MaxAge`, time.Hour*24*30)
	parser.SetDefault(`SystemLogger.RotationTime`, time.Hour*24)
	parser.SetConfigType(`yaml`)
	parser.SetConfigName(`conf`)
}

func parseSystemConfig() {
	parser.AddConfigPath(`./config`)
	if err := parser.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := parser.Unmarshal(SystemConf); err != nil {
		panic(err)
		//} else if err = syscall.Setenv(`TMPDIR`, SystemConf.Http.CacheDir); err != nil {
		//	panic(err)
		//} else if err = initCacheDir(Conf.Http.CacheDir); err != nil {
		//	panic(err)
	} else if err = validate.Struct(SystemConf); err != nil {
		panic(validate.TransError(err, validate.ZH))
	}
}

// IsDebug 系统状态
func IsDebug() bool {
	return SystemConf.SystemStatus == Debug
}

func SetStatus(status string) {
	stat := strings.ToLower(status)
	if stat != Debug && stat != Release {
		return
	}

	SystemConf.SystemStatus = stat
}
