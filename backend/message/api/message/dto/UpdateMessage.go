package dto

import "github.com/victorlazzaroli/microservicesTest/auth/api/message/model"

type UpdateMessageRequest struct {
	model.Message
	Tags []string `json:"tags"`
}
