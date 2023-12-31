package board

import "time"

type Board struct {
	ID          int `gorm:"primaryKey" json:"id"`
	Name        string
	Link        string
	Description string
	Threads     []Thread 
	CreatedAt   time.Time `json:"created_at"`
}

type Thread struct {
	ID   int `gorm:"primaryKey" json:"id"`
	Link string

	Posts []Post
}

type Post struct {
	ID      int `gorm:"primaryKey" json:"id"`
	Link    string
	Content string
	Author  string
	Replies []Post
}
