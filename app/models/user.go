package models

import "strconv"

// User : basic user info
type User struct {
	ID
	GId     string `json:"gid"`
	Email   string `json:"email"`
	Name    string `json:"name" gorm:"not null;comment:User Name"`
	Picture string `json:"picture"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
