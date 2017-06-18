package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	receivers = append(receivers, web)
}

func web() {
	logger := log.New(os.Stdout, "[web] ", log.Lshortfile|log.LstdFlags)
	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("$PORT must be set")
	}
	gin.SetMode("debug")
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.POST("/testtcpmail", func(c *gin.Context) {
		c.Request.ParseForm()
		logger.Println(c.Request.Form.Get("user_name") + " calls " + c.Request.Form.Get("command") + " with: " + c.Request.Form.Get("text"))
	})
	router.Run(":" + port)
}
