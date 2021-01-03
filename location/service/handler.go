package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getLocationByID(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	fmt.Printf("ids: %v; names: %v", ids, names)
	c.String(http.StatusOK, "Hello")
}
