package main

import (
	"net/http"

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
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.Run(":5000")
}

/*
func listener() {
	router := gin.New()
	router.Use(gin.Logger())
	router.GET
	router.Run(":5001")
}*/
