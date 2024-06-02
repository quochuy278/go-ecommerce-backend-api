package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.GET("ping", Pong)
		v1.POST("ping", Pong)
		v1.PATCH("ping", Pong)
		v1.DELETE("ping", Pong)
		v1.HEAD("ping", Pong)
	}

	v2 := r.Group("v2")
	{
		v2.GET("ping", Pong)
		v2.POST("ping", Pong)
		v2.PATCH("ping", Pong)
		v2.DELETE("ping", Pong)
		v2.HEAD("ping", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	// Go support default query

	// Get Param
	// param := c.Param("name") //For /ping:name

	// Get uid
	// uid := c.Query("uid") //For /ping:name
	c.JSON(http.StatusOK, gin.H{
		"message": "ping......pong",
	})
}
