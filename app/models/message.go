package models

// MessageContent define contents of message
type MessageContent struct {
	ContentType string `json:"content_type"`
	Parts       string `json:"parts"`
}

// Message define message
type Message struct {
	ID
	Role    string         `json:"role"`
	ChatID  int            `json:"c_id"`
	Chat    Chat           `json:"chat" gorm:"foreignKey:ChatID;references:ID"`
	Content MessageContent `json:"content" gorm:"type:jsonb"`
	Timestamps
	SoftDeletes
}
