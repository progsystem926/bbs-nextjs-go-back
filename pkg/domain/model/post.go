package model

type Post struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Text      string `json:"text"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
