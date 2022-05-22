package models

import "github.com/google/uuid"

// Book struct = Object of type Book
type Book struct {
	ID              uuid.UUID `json:"id" gorm:"primary_key"`
	Title           string    `json:"title" gorm:"not null"`
	Author          string    `json:"author" gorm:"not null"`
	AvailableCopies int       `json:"available_copies" gorm:"not null"`
	Price           int       `json:"price" gorm:"not null"`
}

type BookDetail struct {
	Title           string `json:"title" `
	Author          string `json:"author"`
	AvailableCopies int    `json:"available_copies" `
	Price           int    `json:"price"`
}

// Declaring some constant variables
const (
	GotAllBooks           = "successfully got all books from database"
	SuccessInsert         = "successfully inserted the book"
	SuccessBookFound      = "successfully found the book"
	SuccessUpdate         = "successfully updated the book record"
	SuccessDelete         = "successfully deleted the book"
	EmptyQuery            = "the context query cannot be empty"
	ErrInsertingBooks     = "failed to insert the book details"
	ErrGettingBooks       = "failed to get books data"
	ErrDeletingBooks      = "failed to delete the book"
	ErrUpdatingBookPrice  = "failed to update book price"
	ErrUpdatingBookCopies = "failed to update book available quantity"
	ErrMatchingBook       = "failed to get the book that has the name provided"
	ErrSettingBooks       = "failed to set books"
	ErrBindingInput       = "failed to bind the context input"
)
