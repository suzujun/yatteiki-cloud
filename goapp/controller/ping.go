package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/suzujun/yatteiki-cloud/goapp/dao"
)

type Ping struct{}

func (p Ping) Get(c *gin.Context) {
	c.JSON(200, gin.H{"message": true})
}

func (p Ping) GetDb(c *gin.Context) {
	c.JSON(200, gin.H{"message": dao.PingDb()})
}
