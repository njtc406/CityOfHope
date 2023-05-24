// Package account
// 模块名:
// 模块功能简介:
package account

import (
	"github.com/njtc406/chaosutil/engine"
	"github.com/njtc406/cityOfHope/internal/pkg/services"
)

// service 账号服务
type service struct {
}

func NewService() engine.IEngine {
	return nil
}

func init() {
	services.Register(`login_service`, NewService())
}
