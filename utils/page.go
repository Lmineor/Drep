package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// ParsePaginateParams return pageNum and pageSize in the Gin.Context
func ParsePaginateParams(c *gin.Context) (int, int) {
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 10
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	return pageNum, pageSize
}
