package models

// ChatListener 表示聊天服务器的侦听器
type ChatListener interface {
	ClientAdded(client Client)     // 添加 client
	ClientRemoved(clientID string) // 删除 client
	MessageSent(message ChatEvent) // 发送消息
}
