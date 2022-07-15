package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Param struct {
	Name string `json:"name"`
}

func LoginPost(c *gin.Context)  {
	var p Param
	if err := c.ShouldBindJSON(&p); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "we need a name",
		})
		return
	}
	 c.JSON(http.StatusOK,gin.H{
		 "message": "Liwenzhou",
	 })

}
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login", LoginPost)
	return router
}