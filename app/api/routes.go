package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {

	Help2()

	root := r.Group("/service", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "the service handler",
		})
	})

	root.GET("static", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "estatic",
		})
	}).Static("static", "./some/static")

	root.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	root.Group("/sub").
		GET("/subep", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "the sub",
			})
		})
}
