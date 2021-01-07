package models

// PingMessage 表示客户端到服务端之间的 ping 的消息
type PingMessage struct {
	MessageType string `json:"type"` // = "PING"
}
