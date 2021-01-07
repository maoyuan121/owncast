package models

import "github.com/owncast/owncast/utils"

// 系统状态
type Status struct {
	Online                bool `json:"online"`      // 是否在线
	ViewerCount           int  `json:"viewerCount"` // 观看者数量
	OverallMaxViewerCount int  `json:"overallMaxViewerCount"`
	SessionMaxViewerCount int  `json:"sessionMaxViewerCount"`

	LastConnectTime    utils.NullTime `json:"lastConnectTime"`    // 最近一次连接时间
	LastDisconnectTime utils.NullTime `json:"lastDisconnectTime"` // 最近一次断开连接时间

	VersionNumber string `json:"versionNumber"` // 系统版本
}
