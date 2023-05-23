package public

import (
	"path"
	"time"
)

const (
	Debug   = `debug`
	Release = `release`
)

type conf struct {
	SystemStatus string        `binding:"oneof=debug release"` // 系统状态(debug/release)
	ProjectName  string        `binding:"required"`            // 项目名称
	CachePath    string        `binding:"required"`            // 系统缓存目录(默认./run)
	LogPath      string        `binding:"omitempty"`           // 日志存储目录(默认./run/logs)
	SystemLogger *SystemLogger `binding:"required"`            // 系统日志
}

type Logger struct {
	Name         string        `binding:""`                // 日志文件名称
	Level        uint32        `binding:"min=0,max=6"`     // 日志写入级别 小于设置级别的类型都会被记录（0:panic/1:fatal/2:error/3:warn/4:info/5:debug/6:trace）
	Caller       bool          `binding:""`                // 是否打印调用者
	Color        bool          `binding:""`                // 是否打印级别色彩
	MaxAge       time.Duration `binding:"min=1m,max=720h"` // 日志保留时间 min=1m,max=720h 最小1分钟,最大1个月
	RotationTime time.Duration `binding:"min=1m,max=24h"`  // 日志切割时间 min=1m,max=24h 最小1分钟,最大1天
}

type RuntimeLogger = Logger
type SystemLogger = Logger

func (s *conf) GetSystemLoggerFileName() string {
	proName := s.ProjectName
	if proName == "" {
		return s.SystemLogger.Name
	}

	if s.SystemLogger.Name == "" {
		return ""
	}

	return path.Join(s.LogPath, s.ProjectName+"_"+s.SystemLogger.Name)
}

type RpcConf struct {
	PoolSize   uint32 `binding:"min=1"` // rpc客户端连接池初始大小(合理配置大小可以降低map重新分配的消耗)
	FailMode   int    `binding:""`      // rpc客户端调用rpc服务器服务失败时的操作方式
	SelectMode int    `binding:""`      // 客户端调用服务时,候选服务的选择模式
	Addr       string `binding:""`      // rpc服务器监听地址
}

type MasterConf struct {
	Addr        string
	ServiceName string
}
