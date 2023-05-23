// Package services
// 模块名: 模块名
// 功能描述: 描述
// 作者:  yr  2023/5/15 0015 22:15
// 最后更新:  yr  2023/5/15 0015 22:15
package services

import "github.com/njtc406/chaosutil/engine"

var Services = make(map[string]engine.IEngine)

// Register 注册服务(由于都是在init或者其他指定地方注册,基本是个单线程过程,所以就不加锁了)
func Register(name string, engine engine.IEngine) {
	Services[name] = engine
}

func Unregister(name string) {
	delete(Services, name)
}
