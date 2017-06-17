package main

import (
	"log"
	"net/http"
	"os"

	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		web()
	}()
	wg.Wait()
}

func web() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.POST("/", func(c *gin.Context) {
		v := c.Request.Form.Get("token")
		log.Println(v)
	})
	router.Run(":" + port)
}
