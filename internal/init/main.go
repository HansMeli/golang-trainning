package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/golang-trainning/internal/routes"
)

func main() {
	r := gin.Default()
	routes.Routes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
