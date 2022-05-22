package handlers

import (
	"learnWednesday/models"
	"learnWednesday/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setting up a interface
type LearnHandler interface {
	GetAllBooks(c *gin.Context)
	InsertABook(c *gin.Context)
	GetBook(c *gin.Context)
	UpdateABookPrice(c *gin.Context)
	UpdateAvailableCopies(c *gin.Context)
	DeleteABook(c *gin.Context)
}

type learnHandler struct {
	service service.LearnWednesdayService
}

func NewServiceHandler(service service.LearnWednesdayService) LearnHandler {
	return &learnHandler{
		service: service,
	}
}

func (svr *learnHandler) GetAllBooks(c *gin.Context) {
	response, err := svr.service.GetAllBooks(c)

	// If the error is not nil
	if err != nil {
		handlerError(c, err.Error())
		// Return at this point
		return
	}

	// If no error and successfully got the response
	handleSuccess(c, response, models.GotAllBooks)
}

func (svr *learnHandler) InsertABook(c *gin.Context) {
	// Validating the input recieved
	var input models.BookDetail
	if err := c.ShouldBindJSON(&input); err != nil {
		handlerError(c, models.ErrBindingInput)
		return
	}

	err := svr.service.InsertABook(c, input)
	if err != nil {
		handlerError(c, err.Error())
		return
	}

	// Returning success message
	var data []interface{}
	data = append(data, input)
	handleSuccess(c, data, models.SuccessInsert)
}

func (svr *learnHandler) GetBook(c *gin.Context) {
	// Reading params from the context
	bookName := c.Query("name")
	// If the bookName in context is empty return error
	if bookName == "" {
		handlerError(c, models.EmptyQuery)
		return
	}

	response, err := svr.service.GetBook(c, bookName)
	// If the error is not nil
	if err != nil {
		handlerError(c, err.Error())
		// Return at this point
		return
	}

	// If no error and successfully got the response
	handleSuccess(c, response, models.SuccessBookFound)
}

func (svr *learnHandler) UpdateABookPrice(c *gin.Context) {
	// Getting id from context
	id := c.Query("id")
	if id == "" {
		handlerError(c, models.EmptyQuery)
		return
	}
	price := c.Query("newPrice")
	if price == "" {
		handlerError(c, models.EmptyQuery)
		return
	}

	err := svr.service.UpdateABookPrice(c, id, price)
	// If the error is not nil
	if err != nil {
		handlerError(c, err.Error())
		// Return at this point
		return
	}

	// If no error and successfully got the response
	handleSuccess(c, nil, models.SuccessUpdate)
}

func (svr *learnHandler) UpdateAvailableCopies(c *gin.Context) {
	// Getting id from context
	id := c.Query("id")
	if id == "" {
		handlerError(c, models.EmptyQuery)
		return
	}
	copies := c.Query("availableCopies")
	if copies == "" {
		handlerError(c, models.EmptyQuery)
		return
	}

	err := svr.service.UpdateAvailableCopies(c, id, copies)
	// If the error is not nil
	if err != nil {
		handlerError(c, err.Error())
		// Return at this point
		return
	}

	// If no error and successfully got the response
	handleSuccess(c, nil, models.SuccessUpdate)
}

func (svr *learnHandler) DeleteABook(c *gin.Context) {
	// Getting id of book to be deleted
	id := c.Query("id")
	if id == "" {
		handlerError(c, models.EmptyQuery)
		return
	}

	err := svr.service.DeleteABook(c, id)
	// If the error is not nil
	if err != nil {
		handlerError(c, err.Error())
		// Return at this point
		return
	}

	// If no error and successfully got the response
	handleSuccess(c, nil, models.SuccessDelete)
}

// function to handle errror
func handlerError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status-code": http.StatusBadRequest,
		"message":     message,
		"error":       true,
	})
}

// function to handle success message
func handleSuccess(c *gin.Context, data []interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status-code": http.StatusOK,
		"message":     message,
		"data":        data,
		"error":       false,
	})
}
