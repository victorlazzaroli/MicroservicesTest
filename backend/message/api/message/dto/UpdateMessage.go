package dto

import "github.com/victorlazzaroli/microservicesTest/message/api/message/model"

type UpdateMessageRequest struct {
	model.Message
	Tags []string `json:"tags"`
}
