// Package works
// 模块名: 模块名
// 功能描述: 描述
// 作者:  yr  2023/5/11 0011 20:56
// 最后更新:  yr  2023/5/11 0011 20:56
package works

import (
	"github.com/njtc406/cityOfHope/web/app/login_server/define"
)

// Login 登录验证
func Login(req *define.ReqLogin, res *define.Response) {
	//retData := &define.ResLoginInfo{
	//	UserName:  req.UserName,
	//	ExtraInfo: req.ExtraInfo,
	//}
	//userInfo, ok := account.Check(req.UserName, req.Password)
	//if !ok {
	//	res.SetStatus(http.StatusNotFound).SetMessage("账号或者密码错误")
	//	return
	//} else {
	//	token, timeStamp := account.CreateToken(userInfo.ID)
	//	retData.UserID = userInfo.ID
	//	retData.ChannelID = userInfo.ChannelID
	//	retData.IsNew = userInfo.IsNew
	//	retData.UserGroup = userInfo.UserGroup
	//	retData.LoginTime = timeStamp
	//	retData.Token = token
	//	retData.Version = `1.0`
	//	retData.Time = chaostime.Now().Unix()
	//	retData.TimeZoneName, retData.TimeZoneOffset = chaostime.Now().Zone()
	//
	//	res.SetData(retData)
	//}

	return
}

// CreateUser 创建账号
func CreateUser(req *define.ReqCreateUser, res *define.Response) {
	//retData := &define.ResLoginInfo{
	//	UserID:    "",
	//	LoginTime: 0,
	//	Token:     "",
	//}
	//
	//userInfo, err := account.Add(req)
	//if err != nil {
	//	res.SetStatus(http.StatusBadRequest).SetMessage(err.Error())
	//	return
	//}
	//
	//token, timeStamp := account.CreateToken(userInfo.ID)
	//retData.UserID = userInfo.ID
	//retData.ChannelID = userInfo.ChannelID
	//retData.IsNew = userInfo.IsNew
	//retData.UserGroup = userInfo.UserGroup
	//retData.LoginTime = timeStamp
	//retData.Token = token
	//retData.Version = `1.0`
	//retData.Time = chaostime.Now().Unix()
	//retData.TimeZoneName, retData.TimeZoneOffset = chaostime.Now().Zone()
	//
	//res.SetData(retData)
}

// ServerList 获取服务器列表
func ServerList(_ *define.ReqServerList, res *define.Response) {

}

// Notice 获取公告
func Notice(req *define.ReqNotice, res *define.Response) {

}
