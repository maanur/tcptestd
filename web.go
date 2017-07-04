package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	receivers = append(receivers, web)
}

func web(output io.Writer) {
	logger := log.New(output, "[web] ", log.Lshortfile|log.LstdFlags)
	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("$PORT must be set")
	}
	gin.SetMode("release")
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"logstr": logbuf,
		})
	})

	router.POST("/test", func(c *gin.Context) {
		c.Request.ParseForm()
		logger.Println(c.Request.Form.Get("user_name") + " calls " + c.Request.Form.Get("command") + " with: " + c.Request.Form.Get("text"))
		queries <- c.Request.Context()
	})
	router.Run(":" + port)
}
