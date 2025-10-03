package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"name": "leedonggyu"})
}
