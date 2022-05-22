package handlers

import (
	"learnWednesday/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(service service.LearnWednesdayService) {
	// Setting up a router for API endpoints
	router := gin.Default()

	// List of API endpoints
	// This is default to check if the router is working or not6
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hi. Welcome to this learning application! Today we will perform some basic CRUD opeartions."})
	})

	serviceHandler := NewServiceHandler(service)

	// Book Library endpoints
	router.Group("/bookLibrary")
	{
		v1Group := router.Group("/bookLibrary/v1")
		{
			v1Group.GET("/bookLibrary", serviceHandler.GetAllBooks)
			v1Group.POST("/insertBook", serviceHandler.InsertABook)
			v1Group.GET("/getBook", serviceHandler.GetBook)
			v1Group.POST("/updateBookPrice", serviceHandler.UpdateABookPrice)
			v1Group.POST("/updateAvailableCopies", serviceHandler.UpdateAvailableCopies)
			v1Group.POST("/deleteBook", serviceHandler.DeleteABook)
		}
	}

	//Running the router
	router.Run("localhost:2022")
}
