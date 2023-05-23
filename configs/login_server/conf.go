package config

import (
	"github.com/njtc406/cityOfHope/configs/public"
	"time"
)

type systemConf struct {
	NodeType         string                `binding:"required"` // 系统类型
	NodeID           uint32                `binding:""`         // 节点id
	ServiceNames     []string              `binding:""`         // 节点提供的服务名称
	ProfilerInterval time.Duration         `binding:"required"` // 性能监控报告间隔(毫秒)如果设置为0表示不开启性能监控
	RuntimeLogger    *public.RuntimeLogger `binding:""`         // 运行日志
	HttpConf         *httpConf             `binding:"required"` // http配置
	RpcConf          *public.RpcConf       `binding:"required"` // rpc配置
	Sources          *sources              `binding:""`         // 资源配置
}

type httpConf struct {
	Addr         string        `binding:"required"` // 监听地址(127.0.0.1:80|0.0.0.0:80)
	ExternalAddr string        `binding:"required"` // 外网监听地址(用于docker等需要端口映射的地方)
	ReadTimeout  time.Duration `binding:""`         // 读超时
	WriteTimeout time.Duration `binding:""`         // 写超时
	CaFileList   []string      `binding:""`         // 证书地址
	Secret       string        `binding:""`         // 校验码
}

type sources struct {
	LevelDB string `binding:""` // leveldb存储路径
	Sqlite  string `binding:""` // sqlite存储路径
}
