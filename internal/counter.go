package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var counter = 0

func Counter(c *gin.Context) {

	c.HTML(http.StatusOK, "counter.html", gin.H{"counter": counter})
}

func Increment(c *gin.Context) {
	counter++
	c.String(http.StatusOK, fmt.Sprintf("Counter : %d", counter))
}

func Decrement(c *gin.Context) {
	counter--
	c.String(http.StatusOK, fmt.Sprintf("Counter : %d", counter))
}
