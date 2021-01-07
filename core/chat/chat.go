package chat

import (
	"errors"
	"time"

	"github.com/owncast/owncast/models"
)

// 设置好一个 chat server
// 1. 设置好数据库
// 2. 实例化一个 server 到 _server
func Setup(listener models.ChatListener) {
	setupPersistence()

	clients := make(map[string]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan models.ChatEvent)
	pingCh := make(chan models.PingMessage)
	doneCh := make(chan bool)
	errCh := make(chan error)

	_server = &server{
		clients,
		"/entry", //hardcoded due to the UI requiring this and it is not configurable
		listener,
		addCh,
		delCh,
		sendAllCh,
		pingCh,
		doneCh,
		errCh,
	}
}

// 开始运行 chat server
// 1. server 开始监听
// 2. server 开始 ping
func Start() error {
	if _server == nil {
		return errors.New("chat server is nil")
	}

	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			_server.ping()
		}
	}()

	_server.Listen()

	return errors.New("chat server failed to start")
}

// 发送消息给所有 client
func SendMessage(message models.ChatEvent) {
	if _server == nil {
		return
	}

	_server.SendToAll(message)
}

// 获取历史消息（一天内）
func GetMessages(filtered bool) []models.ChatEvent {
	if _server == nil {
		return []models.ChatEvent{}
	}

	return getChatHistory(filtered)
}

// 根据 clientid 获取 client
func GetClient(clientID string) *Client {
	for _, client := range _server.Clients {
		if client.ClientID == clientID {
			return client
		}
	}
	return nil
}
