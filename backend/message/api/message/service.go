package message

import (
	"github.com/victorlazzaroli/microservicesTest/auth/api/message/dto"
	"github.com/victorlazzaroli/microservicesTest/auth/api/message/model"
)

type MessageServiceI interface {
	CreateMessage(request *dto.CreateMessageRequest) error
	DeleteMessage(request int) error
	GetMessage(request int) (model.Message, error)
	GetMessages(authorId int) ([]model.Message, error)
	UpdateMessage(id int, request dto.UpdateMessageRequest) error
}

type serviceMessage struct {
	repository MessageRepositoryI
}

// GetMessage implements MessageServiceI.
func (s *serviceMessage) GetMessage(request int) (model.Message, error) {
	return s.repository.GetById(request)
}

// GetMessages implements MessageServiceI.
func (s *serviceMessage) GetMessages(authorId int) ([]model.Message, error) {
	if authorId < 0 {
		return s.repository.GetAll()
	}

	return s.repository.GetAllByAuthor(authorId)
}

func (s *serviceMessage) CreateMessage(request *dto.CreateMessageRequest) error {

	var tags []model.Tag
	for _, tag := range request.Tags {
		newTag := model.Tag{
			Name: tag,
		}

		tags = append(tags, newTag)
	}

	return s.repository.Save(&model.Message{
		Title:    request.Title,
		Text:     request.Text,
		Author:   request.Author,
		AuthorID: request.AuthorID,
		Tags:     tags,
	})
}

func (s *serviceMessage) UpdateMessage(id int, request dto.UpdateMessageRequest) error {
	var tags []model.Tag
	for _, tag := range request.Tags {
		newTag := model.Tag{
			Name: tag,
		}

		tags = append(tags, newTag)
	}

	modelMessage := model.Message{
		Title:    request.Title,
		Text:     request.Text,
		Author:   request.Author,
		AuthorID: request.AuthorID,
		Tags:     tags,
	}

	modelMessage.ID = uint(id)

	return s.repository.Update(&modelMessage)
}

func (s *serviceMessage) DeleteMessage(request int) error {
	return s.repository.Delete(request)
}

func NewMessageService(repository MessageRepositoryI) MessageServiceI {
	return &serviceMessage{
		repository: repository,
	}
}
