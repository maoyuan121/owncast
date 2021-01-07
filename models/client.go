package models

import (
	"net/http"
	"time"

	"github.com/owncast/owncast/geoip"
	"github.com/owncast/owncast/utils"
)

type ConnectedClientsResponse struct {
	Clients []Client `json:"clients"`
}

type Client struct {
	ConnectedAt  time.Time         `json:"connectedAt"`  // 连接时间
	LastSeen     time.Time         `json:"-"`            // 最后一次连接时间？？？
	MessageCount int               `json:"messageCount"` // 消息数
	UserAgent    string            `json:"userAgent"`    // 浏览器 useragent
	IPAddress    string            `json:"ipAddress"`    // ip 地址
	Username     *string           `json:"username"`     // 用户名
	ClientID     string            `json:"clientID"`     // client  id
	Geo          *geoip.GeoDetails `json:"geo"`          // geo 信息
}

func GenerateClientFromRequest(req *http.Request) Client {
	return Client{
		ConnectedAt:  time.Now(),
		LastSeen:     time.Now(),
		MessageCount: 0,
		UserAgent:    req.UserAgent(),
		IPAddress:    utils.GetIPAddressFromRequest(req),
		Username:     nil,
		ClientID:     utils.GenerateClientIDFromRequest(req),
	}
}
