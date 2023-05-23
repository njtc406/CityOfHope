package config

import (
	"github.com/njtc406/chaosutil/validate"
	"github.com/njtc406/cityOfHope/configs/public"
	"github.com/spf13/viper"
)

var (
	SystemConf = new(systemConf)
	parser     = viper.New()
)

func init() {
	// TODO 这里应该是需要修改的,一部分系统配置可能是由集群控制节点下发的
	public.InitPublicConf()
	setDefaultConf()
	parseSystemConfig()
}

func setDefaultConf() {
	parser.SetConfigType(`yaml`)
	parser.SetConfigName(`conf`)
}

func parseSystemConfig() {
	parser.AddConfigPath(`./config/login_server`)
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
