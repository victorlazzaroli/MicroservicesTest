package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Title    string `json:"title"`
	Text     string `json:"text"`
	Author   string `json:"author"`
	AuthorID int    `json:"author_id"`
	Likes    uint   `json:"likes" gorm:"default:0"`
	Dislikes uint   `json:"dislikes" gorm:"default:0"`
	Tags     []Tag  `gorm:"many2many:message_tags"`
}
