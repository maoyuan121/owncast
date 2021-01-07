package chat

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"

	"github.com/owncast/owncast/config"
	"github.com/owncast/owncast/models"
)

var (
	_server *server
)

var l = sync.Mutex{}

// Server represents the server which handles the chat.
type server struct {
	Clients map[string]*Client // 连接到的客户端集合

	pattern  string              // url pattern
	listener models.ChatListener // 聊天服务器的侦听器

	addCh     chan *Client            // 添加客户端的通道
	delCh     chan *Client            // 删除客户端的通道
	sendAllCh chan models.ChatEvent   // 发送消息的通道
	pingCh    chan models.PingMessage // ping 通道
	doneCh    chan bool               //
	errCh     chan error              // 错误通道
}

// 写 add client 通道
func (s *server) add(c *Client) {
	s.addCh <- c
}

// 写 del client 通道
func (s *server) remove(c *Client) {
	s.delCh <- c
}

// SendToAll 发送一个消息给所有的客户端
func (s *server) SendToAll(msg models.ChatEvent) {
	s.sendAllCh <- msg
}

// 写 err 通道
func (s *server) err(err error) {
	s.errCh <- err
}

// 发送消息给所有 client
func (s *server) sendAll(msg models.ChatEvent) {
	for _, c := range s.Clients {
		c.write(msg)
	}
}

// 写所有 client 的 ping 通道
func (s *server) ping() {
	ping := models.PingMessage{MessageType: PING}
	for _, c := range s.Clients {
		c.pingch <- ping
	}
}

// 写所有 client 的 usernameChange 通道
func (s *server) usernameChanged(msg models.NameChangeEvent) {
	for _, c := range s.Clients {
		c.usernameChangeChannel <- msg
	}
}

// 建立连接
// 创建一个 client 并添加到 server 的 add client 通道
// client 开始侦听
func (s *server) onConnection(ws *websocket.Conn) {
	client := NewClient(ws)

	defer func() {
		s.removeClient(client)

		if err := ws.Close(); err != nil {
			s.errCh <- err
		}
	}()

	s.add(client)
	client.listen()
}

// 侦听
// 处理 client 连接和广播请求
func (s *server) Listen() {
	http.Handle(s.pattern, websocket.Handler(s.onConnection))

	log.Tracef("Starting the websocket listener on: %s", s.pattern)

	for {
		select {
		// 从 add client 通道读
		// 添加到 server 的 client 集合
		// 发送欢迎消息给这个 client
		//
		case c := <-s.addCh:
			l.Lock()
			s.Clients[c.socketID] = c
			l.Unlock()

			s.listener.ClientAdded(c.GetViewerClientFromChatClient())
			s.sendWelcomeMessageToClient(c)

		// 从 del client 通道读
		case c := <-s.delCh:
			s.removeClient(c)
			// 从发送消息通道读
		case msg := <-s.sendAllCh:
			// message was received from a client and should be sanitized, validated
			// and distributed to other clients.
			//
			// Will turn markdown into html, sanitize user-supplied raw html
			// and standardize this message into something safe we can send everyone else.
			msg.RenderAndSanitizeMessageBody()

			s.listener.MessageSent(msg)
			s.sendAll(msg)

			// 将消息保存到数据库中
			addMessage(msg)
			// 从 ping 通道读
		case ping := <-s.pingCh:
			fmt.Println("PING?", ping)
			// 从 err 通道读
		case err := <-s.errCh:
			log.Error("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}

// 删除一个 client
// 从 client 集合中删除这个 client
//
func (s *server) removeClient(c *Client) {
	l.Lock()

	if _, ok := s.Clients[c.socketID]; ok {
		delete(s.Clients, c.socketID)

		s.listener.ClientRemoved(c.socketID)

		log.Tracef("The client was connected for %s and sent %d messages (%s)", time.Since(c.ConnectedAt), c.MessageCount, c.ClientID)
	}
	l.Unlock()
}

// 发送欢迎消息给 client
// client 刚连接上来的时候发送
func (s *server) sendWelcomeMessageToClient(c *Client) {
	go func() {
		// Add an artificial delay so people notice this message come in.
		time.Sleep(7 * time.Second)

		initialChatMessageText := fmt.Sprintf("Welcome to %s! %s", config.Config.InstanceDetails.Title, config.Config.InstanceDetails.Summary)
		initialMessage := models.ChatEvent{ClientID: "owncast-server", Author: config.Config.InstanceDetails.Name, Body: initialChatMessageText, ID: "initial-message-1", MessageType: "SYSTEM", Visible: true, Timestamp: time.Now()}
		c.write(initialMessage)
	}()
}
