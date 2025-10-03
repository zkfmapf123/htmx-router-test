package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zkfmapf123/htmx-test/internal"
)

func main() {
	g := gin.Default()

	// load templates
	g.LoadHTMLGlob("./templates/*.html")

	// router
	g.GET("/health", internal.RouterHealth)

	counterGroup := g.Group("/counter")
	{
		counterGroup.GET("/", internal.RouterCounter)
		counterGroup.POST("/increment", internal.Increment)
		counterGroup.POST("/decrease", internal.Decrease)
	}

	g.Run(":8080")

}
