package dev

import (
	"github.com/drep/source"
	"github.com/gin-gonic/gin"
)

func Pop(c *gin.Context) {
	populateDb()
}

func populateDb() {
	source.Casbin.Init()
}
