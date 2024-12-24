package dto

import "gorm.io/gorm"

type CreateMessageRequest struct {
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	Author   string   `json:"author"`
	AuthorID int      `json:"author_id"`
	Tags     []string `json:"tags"`
}

type CreateMessageResponse struct {
	gorm.Model
	CreateMessageRequest
	Likes    uint `json:"likes"`
	Dislikes uint `json:"dislikes"`
}
