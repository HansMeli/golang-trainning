package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	comics "github.com/golang-trainning/internal/controllers"
	routes "github.com/golang-trainning/internal/routes"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func updateDatabase(t time.Time) {
	resUpdate := comics.UpdateComicsLibrary()
	fmt.Printf(resUpdate, t)
}

func main() {
	r := gin.Default()
	routes.Routes(r)
	r.GET("/sync_database", func(c *gin.Context) {
		go doEvery(60000*time.Millisecond, updateDatabase)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
