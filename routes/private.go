package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPrivateRoutes(r *gin.Engine) {
	privateRoutes := r.Group("/private")

	privateRoutes.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})
}
