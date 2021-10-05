package main

import (
	"basic-gin/controller"
	"basic-gin/handle"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(Authentication)

	// validate := validator.New()
	// validate.RegisterValidation("is_charan", ValidateCharan)

	m := handle.NewMember()

	r.GET("/member", controller.AllData(m))

	r.POST("/member", controller.AddData(m))

	r.Run()

}

func Authentication(c *gin.Context) {
	// fmt.Println("Authen...")
	token := c.GetHeader("Authorization")

	auth := strings.Split(token, " ")

	if len(auth) != 2 || auth[0] != "Bearer" {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Authorization faild...",
		})
		return
	}

	if auth[1] != "asdfghjkl" {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Invalid API Token...",
		})
		return
	}

	c.Next()
}

// func ValidateCharan(field validator.FieldLevel) bool {
// 	return strings.Contains(field.Field().String(), "charan")
// }
