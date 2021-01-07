package models

// NameChangeEvent 表示用户在聊天室修改了名字
type NameChangeEvent struct {
	OldName string `json:"oldName"` // 原名
	NewName string `json:"newName"` // 修改后的名字
	Image   string `json:"image"`   //
	Type    string `json:"type"`    // NAME_CHANGE
	ID      string `json:"id"`      // 这个消息的唯一 id
}
