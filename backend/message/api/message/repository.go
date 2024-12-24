package message

import (
	"github.com/victorlazzaroli/microservicesTest/auth/api/message/model"
	"gorm.io/gorm"
)

type MessageRepositoryI interface {
	Save(message *model.Message) error
	GetById(id int) (model.Message, error)
	GetAll() ([]model.Message, error)
	GetAllByAuthor(authorID int) ([]model.Message, error)
	Update(message *model.Message) error
	Delete(id int) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) messageRepository {
	var repository messageRepository
	repository = messageRepository{
		db: db,
	}
	return repository
}

func (repository *messageRepository) Save(message *model.Message) error {
	return repository.db.Save(message).Error
}

func (repository *messageRepository) GetById(id int) (model.Message, error) {

	var result model.Message
	errore := repository.db.First(&result, id).Error

	return result, errore
}

func (repository *messageRepository) GetAllByAuthor(authorID int) ([]model.Message, error) {

	result := []model.Message{}
	errore := repository.db.Find(&result, "author_id = ?", authorID).Error

	return result, errore
}

func (repository *messageRepository) GetAll() ([]model.Message, error) {
	result := []model.Message{}
	errore := repository.db.Find(&result).Error

	return result, errore
}

func (repository *messageRepository) Update(message *model.Message) error {
	return repository.db.Save(message).Error
}

func (repository *messageRepository) Delete(id int) error {
	var message model.Message
	err := repository.db.Delete(&message, id).Error
	return err
}
