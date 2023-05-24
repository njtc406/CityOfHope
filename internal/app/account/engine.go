// Package account
// 模块名: 账号服务
// 模块功能简介: 用于提供账号相关服务
package account

import (
	"github.com/njtc406/chaosutil/engine"
	"github.com/njtc406/cityOfHope/internal/pkg/services"
)

// service 账号服务
type service struct {
	// TODO: 账号服务只需要监听rpc消息,然后提供创建和校验接口
	// TODO: 然后做一个查询服务,也使用rpc,供后台查询使用
	// TODO: 数据库直接使用mysql来做
}

func (s *service) Init() {

}

func (s *service) Start() {

}

func (s *service) Stop() {

}

func NewService() engine.IEngine {
	return &service{}
}

func init() {
	services.Register(`login_service`, NewService())
}
