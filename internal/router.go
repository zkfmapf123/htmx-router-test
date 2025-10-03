package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterHealth(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"message": "OK"})
	c.HTML(http.StatusOK, "index.html", gin.H{"name": "leedonggyu"})
}

var count = 0

func RouterCounter(c *gin.Context) {
	c.HTML(http.StatusOK, "counter.html", gin.H{"count": count})
}

func Increment(c *gin.Context) {

	count++
	c.String(http.StatusOK, strconv.Itoa(count))
}

func Decrease(c *gin.Context) {
	count--
	c.String(http.StatusOK, strconv.Itoa(count))
}
