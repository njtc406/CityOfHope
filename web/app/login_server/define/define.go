// Package define
// 模块名: 模块名
// 功能描述: 描述
// 作者:  yr  2023/5/11 0011 20:09
// 最后更新:  yr  2023/5/11 0011 20:09
package define

import "net/http"

type Response struct {
	Code    int         `json:"code"`           // 返回码
	Message string      `json:"message"`        // 返回消息
	Data    interface{} `json:"data,omitempty"` // 返回数据
}

func NewResponse() *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: "ok",
	}
}

func (r *Response) SetStatus(code int) *Response {
	r.Code = code
	return r
}
func (r *Response) SetMessage(msg string) *Response {
	r.Message = msg
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) Success() bool {
	return r.Code == http.StatusOK
}

// ReqLogin 登录请求
type ReqLogin struct {
	UserName  string      `json:"user_name" binding:"required,min=6,max=18"` // 用户名
	Password  string      `json:"password" binding:"required,min=5,max=64"`  // 密码
	ExtraInfo interface{} `json:"extra_info" binding:""`                     // 客户端透传数据
}

// ResLoginInfo 登录基础回复响应
type ResLoginInfo struct {
	UserID         string      `json:"user_id"`          // 账号id
	UserName       string      `json:"user_name"`        // 账号
	ChannelID      string      `json:"channel_id"`       // 渠道id
	IsNew          bool        `json:"is_new"`           // 是否是新角色
	ExtraInfo      interface{} `json:"extra_info"`       // 客户端透传信息
	UserGroup      uint8       `json:"user_group"`       // 账号分类(1白名单账号)
	LoginTime      int64       `json:"login_time"`       // 登入时间
	Token          string      `json:"token"`            // 令牌
	Version        string      `json:"version"`          // 服务器版本(这个预留给之后做更新什么的)
	TimeZoneName   string      `json:"time_zone"`        // 服务器时区
	TimeZoneOffset int         `json:"time_zone_offset"` // 服务器时区时间偏移
	Time           int64       `json:"time"`             // 服务器当前时间
}

// ReqCreateUser 创建账号请求
type ReqCreateUser struct {
	UserName       string `json:"user_name"`                  // 用户名
	Password       string `json:"password"`                   // 密码
	CardID         string `json:"card_id"`                    // 身份证号
	Name           string `json:"name"`                       // 真实姓名
	Phone          string `json:"phone"`                      // 手机号
	Email          string `json:"email,omitempty"`            // 邮箱
	ChannelID      string `json:"channel_id"`                 // 渠道id
	ChildChannelID string `json:"child_channel_id,omitempty"` // 子渠道id
	PhoneModel     string `json:"phone_model,omitempty"`      // 手机型号
	SystemVersion  string `json:"system_version,omitempty"`   // 手机系统版本
}

// ReqServerList 服务器列表请求
type ReqServerList struct {
}

type ServerInfo struct {
	ServerID   string  `json:"server_id"`   // 服务器id
	ShowID     string  `json:"show_id"`     // 显示id
	Name       string  `json:"name"`        // 服务器名称
	Status     uint8   `json:"status"`      // 服务器状态
	CreateTime int64   `json:"create_time"` // 服务器创建时间
	OpenTime   int64   `json:"open_time"`   // 服务器开服时间
	ShowTime   int64   `json:"show_time"`   // 服务器对外开放时间(设置了自动开服的服务器才有)
	Tips       *string `json:"tips"`        // 提示信息
}

// ResServerList 服务器列表响应
type ResServerList struct {
	Response
	Data []*ServerInfo `json:"data"`
}

// ReqNotice 公告请求
type ReqNotice struct {
	PackageName string `json:"package_name"` // 包名(这个有点蠢,应该可以换个其他的方式)
}

// NoticeInfo 公告内容
type NoticeInfo struct {
	Title       *string `json:"title"`        // 标题
	SubTitle    *string `json:"sub_title"`    // 副标题
	Content     *string `json:"content"`      // 内容
	Type        *string `json:"type"`         // 类型
	CreateTime  int64   `json:"create_time"`  // 创建时间
	ExpiredTime int64   `json:"expired_time"` // 过期时间
	ShowTime    int64   `json:"show_time"`    // 显示时间
}
