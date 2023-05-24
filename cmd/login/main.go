// Package main
// 模块名:
// 模块功能简介:
package main

import (
	"github.com/njtc406/cityOfHope/internal/pkg/node"
	_ "github.com/njtc406/cityOfHope/web/app/login_server"
)

func main() {
	node.Start()
}
