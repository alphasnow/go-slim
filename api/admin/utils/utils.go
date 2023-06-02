package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// QueryID
// get url id
func QueryID(c *gin.Context) int {
	idStr := c.Param("id")
	if idStr == "" {
		return 0
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0
	}
	return id
}
