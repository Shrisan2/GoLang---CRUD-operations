package service

import (
	"learnWednesday/models"

	"github.com/gin-gonic/gin"
)

type LearnWednesdayService interface {
	GetAllBooks(c *gin.Context) ([]interface{}, error)
	InsertABook(c *gin.Context, bookDetail models.BookDetail) error
	GetBook(c *gin.Context, bookName string) ([]interface{}, error)
	UpdateABookPrice(c *gin.Context, id, price string) error
	UpdateAvailableCopies(c *gin.Context, id, copies string) error
	DeleteABook(c *gin.Context, id string) error
}

type learnWednesdayService struct {
	db Store
}

func NewLearnWednesdayService(store Store) LearnWednesdayService {
	return &learnWednesdayService{
		db: store,
	}
}

// Handler to get all books
func (service learnWednesdayService) GetAllBooks(c *gin.Context) ([]interface{}, error) {
	var responses []interface{}

	response, err := service.db.GetAllBooks(c)
	if err != nil {
		return nil, err
	}

	// Appeding the response to the responses
	responses = append(responses, response)

	return responses, nil
}

func (service learnWednesdayService) InsertABook(c *gin.Context, bookDetail models.BookDetail) error {
	err := service.db.InsertABook(c, bookDetail)
	if err != nil {
		return err
	}
	return nil
}

func (service learnWednesdayService) GetBook(c *gin.Context, bookName string) ([]interface{}, error) {
	var responses []interface{}

	response, err := service.db.GetBook(c, bookName)
	if err != nil {
		return nil, err
	}

	// Appeding the response to the responses
	responses = append(responses, response)

	return responses, nil
}

func (service learnWednesdayService) UpdateABookPrice(c *gin.Context, id, price string) error {
	err := service.db.UpdateABookPrice(c, id, price)
	if err != nil {
		return err
	}
	return nil
}

func (service learnWednesdayService) UpdateAvailableCopies(c *gin.Context, id, copies string) error {
	err := service.db.UpdateAvailableCopies(c, id, copies)
	if err != nil {
		return err
	}
	return nil
}

func (service learnWednesdayService) DeleteABook(c *gin.Context, id string) error {
	err := service.db.DeleteABook(c, id)
	if err != nil {
		return err
	}
	return nil
}
