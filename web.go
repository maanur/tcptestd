package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var web = Web{name: "web"}

func init() {
	backRunList = append(backRunList, &web)
}

// Web handles incoming queries and provides UI. Implements Integrator for consistency.
type Web struct {
	name  string
	logwr *io.Writer
}

func (w *Web) CallName() string {
	return w.name
}

func (w *Web) Run(ctx context.Context, output io.Writer) {
	log := log.New(output, "[web] ", log.Lshortfile|log.LstdFlags)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	gin.SetMode("release")
	gin.DefaultWriter = output
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"logstr": logr.buf.String(),
		})
	})

	router.POST("/test", func(c *gin.Context) {
		c.Request.ParseForm()
		log.Println(c.Request.Form.Get("user_name") + " calls " + c.Request.Form.Get("command") + " with: " + c.Request.Form.Get("text"))
	})
	router.Run(":" + port)
	<-ctx.Done()
}

func (w *Web) HandlerFunc(ctx *gin.Context) {
	ctx.Abort()
}
