package models

// 聊天的 Emoji
type CustomEmoji struct {
	Name  string `json:"name"`  // 文件名
	Emoji string `json:"emoji"` // 文件路径
}
