package service

import (
	"errors"
	"learnWednesday/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store interface {
	GetAllBooks(c *gin.Context) ([]models.Book, error)
	InsertABook(c *gin.Context, bookDetail models.BookDetail) error
	GetBook(c *gin.Context, bookName string) ([]models.Book, error)
	UpdateABookPrice(c *gin.Context, id, price string) error
	UpdateAvailableCopies(c *gin.Context, id, copies string) error
	DeleteABook(c *gin.Context, id string) error
}

type mySQLStore struct {
	db *gorm.DB
}

func NewMySQLStore(db *gorm.DB) Store {
	return &mySQLStore{
		db: db,
	}
}

func (ms *mySQLStore) GetAllBooks(c *gin.Context) ([]models.Book, error) {
	var books []models.Book
	err := ms.db.Table("books").Find(&books).Error
	if err != nil || len(books) == 0 {
		return nil, errors.New(models.ErrGettingBooks)
	}

	return books, nil
}

func (ms *mySQLStore) InsertABook(c *gin.Context, bookDetail models.BookDetail) error {
	// Inserting the book that contains the bookDetails
	book := models.Book{
		ID:              uuid.New(),
		Title:           bookDetail.Title,
		Author:          bookDetail.Author,
		AvailableCopies: bookDetail.AvailableCopies,
		Price:           bookDetail.Price,
	}

	err := ms.db.Create(&book).Error
	if err != nil {
		return errors.New(models.ErrInsertingBooks)
	}

	return nil
}

func (ms *mySQLStore) GetBook(c *gin.Context, bookName string) ([]models.Book, error) {
	var books []models.Book
	err := ms.db.Table("books").Where("title like ?", bookName+"%").Find(&books).Error
	if err != nil || len(books) == 0 {
		return nil, errors.New(models.ErrMatchingBook)
	}

	return books, nil
}

func (ms *mySQLStore) UpdateABookPrice(c *gin.Context, id, price string) error {
	idUuid := uuid.MustParse(id)
	newPrice, _ := strconv.Atoi(price)
	err := ms.db.Table("books").Where("id = ?", idUuid).Updates(map[string]interface{}{"price": newPrice}).Error
	if err != nil {
		return errors.New(models.ErrUpdatingBookPrice)
	}
	return nil
}

func (ms *mySQLStore) UpdateAvailableCopies(c *gin.Context, id, copies string) error {
	idUuid := uuid.MustParse(id)
	newQuantity, _ := strconv.Atoi(copies)
	err := ms.db.Table("books").Where("id = ?", idUuid).Updates(map[string]interface{}{"available_copies": newQuantity}).Error
	if err != nil {
		return errors.New(models.ErrUpdatingBookCopies)
	}
	return nil
}

func (ms *mySQLStore) DeleteABook(c *gin.Context, id string) error {
	idUuid := uuid.MustParse(id)
	var books []models.Book
	err := ms.db.Table("books").Where("id = ?", idUuid).Delete(&books).Error
	if err != nil {
		return errors.New(models.ErrDeletingBooks)
	}
	return nil

}
