package dev

import (
	"github.com/drep/source"
	"github.com/gin-gonic/gin"
)

func InitDb(c *gin.Context) {
	source.Api.Init()
	source.Admin.Init()
	source.Casbin.Init()
	source.BaseMenu.Init()
	source.Authority.Init()
	source.AdminDpMap.Init()
	source.AuthoritiesMenus.Init()
	source.AuthorityMenu.Init()

}
