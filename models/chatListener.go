package models

// ChatListener 表示聊天服务器的侦听器
type ChatListener interface {
	ClientAdded(client Client)
	ClientRemoved(clientID string)
	MessageSent(message ChatEvent)
}
