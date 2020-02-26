package routes

import (
	"github.com/gin-gonic/gin"
	comics "github.com/golang-trainning/internal/controllers"
)

// Routes group and handle the request for the library
func Routes(e *gin.Engine) {
	// Routes for library inventory
	// marvel := e.Group("/marvel")
	// marvel.GET("/getAllComics", comics.GetAllComics)

	// Routes for library inventory
	library := e.Group("/library")
	library.GET("/getAllComics", comics.GetLibraryComics)
}
