// Package pid
// 模块名: 模块名
// 功能描述: 描述
// 作者:  yr  2023/4/22 0022 2:03
// 最后更新:  yr  2023/4/22 0022 2:03
package pid

import (
	"fmt"
	"github.com/njtc406/cityOfHope/configs/public"

	"os"
	"path"
	"strconv"
	"syscall"
)

// RecordPID 记录pid
func RecordPID(proName, nodeType string, nodeID uint32) {
	os.WriteFile(path.Join(public.Conf.CachePath, fmt.Sprintf("%s_%s_%d", proName, nodeType, nodeID)+".pid"), ([]byte)(strconv.Itoa(syscall.Getpid())), 0644)
}

// DelPID 删除pid
func DelPID(proName, nodeType string, nodeID uint32) {
	os.Remove(path.Join(public.Conf.CachePath, fmt.Sprintf("%s_%s_%d", proName, nodeType, nodeID)+".pid"))
}

// GetPID 获取pid
func GetPID() int {
	return syscall.Getpid()
}
