package models

import (
	"github.com/owncast/owncast/utils"
)

// 系统的统计信息
type Stats struct {
	SessionMaxViewerCount int            `json:"sessionMaxViewerCount"`
	OverallMaxViewerCount int            `json:"overallMaxViewerCount"`
	LastDisconnectTime    utils.NullTime `json:"lastDisconnectTime"` // 最近一次断开连接的时间

	StreamConnected bool              `json:"-"`
	LastConnectTime utils.NullTime    `json:"-"` // 最近一次连接的时间
	Clients         map[string]Client `json:"-"` // 所有连接客户端
}
