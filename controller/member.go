package controller

import (
	"basic-gin/handle"
	"basic-gin/model"

	"github.com/gin-gonic/gin"
)

func AllData(m *handle.MemberData) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, m.AllData())
	}
}

func AddData(m *handle.MemberData) func(c *gin.Context) {
	return func(c *gin.Context) {
		var v model.Member
		err := c.ShouldBindJSON(&v)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		m.AddData(v)
		c.JSON(200, v)
	}
}
