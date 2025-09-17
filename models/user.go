ackage models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `gorm:"uniqueIndex;size:100" json:"username"`
	Email     string    `gorm:"uniqueIndex;size:200" json:"email"`
	Password  string    `json:"-"`
	Role      string    `gorm:"size:20" json:"role"` // "user" or "admin"
}
