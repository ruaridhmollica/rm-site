package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("static/templates/*.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/acting", func(c *gin.Context) {
		c.HTML(http.StatusOK, "acting.html", nil)
	})

	router.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.html", nil)
	})

	router.GET("/projects/trail", func(c *gin.Context) {
		c.HTML(http.StatusOK, "trail.html", nil)
	})

	router.GET("/projects/iglu", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iglu.html", nil)
	})

	router.GET("/cv", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cv.html", nil)
	})

	router.Run(":" + port)
}
