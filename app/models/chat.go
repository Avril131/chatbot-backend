package models

// Chat table to store chat info
type Chat struct {
	ID
	Name   string `json:"name" gorm:"not null;comment:Chat Name"`
	UserID int    `json:"u_id" gorm:"not null"`
	User   User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Timestamps
	SoftDeletes
}
