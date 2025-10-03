package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zkfmapf123/htmx-test/internal"
)

func main() {
	g := gin.Default()

	g.Static("/css", "./templates/css")
	g.Static("/js", "./templates/js")
	g.LoadHTMLGlob("./templates/views/*.html")

	// router
	g.GET("/", internal.DefaultHome)

	counterGroup := g.Group("/counter")
	{
		counterGroup.GET("/", internal.Counter)
		counterGroup.POST("/increment", internal.Increment)
		counterGroup.POST("/decrease", internal.Decrement)
	}

	g.Run(":8081")

}
