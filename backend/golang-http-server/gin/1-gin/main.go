package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/message", func(ctx *gin.Context) {
		ctx.JSON(200, Test{Message: "Hello World"})
	})

	r.GET("/say-hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello, %s!", name)
	})

	r.GET("/say-bye", func(ctx *gin.Context) {
		name := ctx.Query("name")
		if name == "" {
			ctx.String(http.StatusOK, "Bye!")
			return
		}
		ctx.String(http.StatusOK, "Bye, %s!", name)
	})

	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		name := c.DefaultPostForm("name", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"name":    name,
		})
	})

	r.LoadHTMLGlob("*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "HTML Render",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/login", loginEndpoint) // /v1/login
		// v1.POST("/submit", submitEndpoint) // /v1/submit
		// v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", loginEndpoint) // /v2/login
		// v2.POST("/submit", submitEndpoint)
		// v2.POST("/read", readEndpoint)
	}

	r.Run()
}

func loginEndpoint(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "roger" || json.Password != "12341234" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

type Login struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

type Test struct {
	Message string `json:"message"`
}
