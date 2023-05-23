package config

import (
	"path"
	"time"
)

const (
	Debug   = `debug`
	Release = `release`
)

type systemConf struct {
	SystemStatus     string         `binding:"oneof=debug release"` // 系统状态(debug/release)
	ProjectName      string         `binding:"required"`            // 项目名称
	NodeType         string         `binding:"required"`            // 系统类型
	NodeID           uint32         `binding:""`                    // 节点id
	ServiceNames     []string       `binding:""`                    // 节点提供的服务名称
	CachePath        string         `binding:"required"`            // 系统缓存目录(默认./run)
	LogPath          string         `binding:"omitempty"`           // 日志存储目录(默认./run/logs)
	ProfilerInterval time.Duration  `binding:"required"`            // 性能监控报告间隔(毫秒)如果设置为0表示不开启性能监控
	SystemLogger     *systemLogger  `binding:"required"`            // 系统日志
	RuntimeLogger    *runtimeLogger `binding:""`                    // 运行日志
	HttpConf         *httpConf      `binding:"required"`            // http配置
	RpcConf          *rpcConf       `binding:"required"`            // rpc配置
	Sources          *sources       `binding:""`                    // 资源配置
}

type httpConf struct {
	Addr         string        `binding:"required"` // 监听地址(127.0.0.1:80|0.0.0.0:80)
	ExternalAddr string        `binding:"required"` // 外网监听地址(用于docker等需要端口映射的地方)
	ReadTimeout  time.Duration `binding:""`         // 读超时
	WriteTimeout time.Duration `binding:""`         // 写超时
	CaFileList   []string      `binding:""`         // 证书地址
	Secret       string        `binding:""`         // 校验码
}

type logger struct {
	Name         string        `binding:""`                // 日志文件名称
	Level        uint32        `binding:"min=0,max=6"`     // 日志写入级别 小于设置级别的类型都会被记录（0:panic/1:fatal/2:error/3:warn/4:info/5:debug/6:trace）
	Caller       bool          `binding:""`                // 是否打印调用者
	Color        bool          `binding:""`                // 是否打印级别色彩
	MaxAge       time.Duration `binding:"min=1m,max=720h"` // 日志保留时间 min=1m,max=720h 最小1分钟,最大1个月
	RotationTime time.Duration `binding:"min=1m,max=24h"`  // 日志切割时间 min=1m,max=24h 最小1分钟,最大1天
}

type runtimeLogger = logger
type systemLogger = logger

func (s *systemConf) GetSystemLoggerFileName() string {
	proName := s.ProjectName
	if proName == "" {
		return s.SystemLogger.Name
	}

	if s.SystemLogger.Name == "" {
		return ""
	}

	return path.Join(s.LogPath, s.ProjectName+"_"+s.SystemLogger.Name)
}

type rpcConf struct {
	PoolSize   uint32 `binding:"min=1"` // rpc客户端连接池初始大小(合理配置大小可以降低map重新分配的消耗)
	FailMode   int    `binding:""`      // rpc客户端调用rpc服务器服务失败时的操作方式
	SelectMode int    `binding:""`      // 客户端调用服务时,候选服务的选择模式
	Addr       string `binding:""`      // rpc服务器监听地址
}

type masterConf struct {
	Addr        string
	ServiceName string
}

type sources struct {
	LevelDB string `binding:""` // leveldb存储路径
	Sqlite  string `binding:""` // sqlite存储路径
}
