package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

	app.Static("/static", "./static")

	app.LoadHTMLGlob("templates/*")

	app.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", gin.H{})
	})

	app.Run(":8080")
}
