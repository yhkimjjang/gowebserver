package routes

import (
	"net/http"

	"yhkim/gowebserver/services"

	"github.com/gin-gonic/gin"
)

func addSearchRoute(rg *gin.RouterGroup) {
	ping := rg.Group("/search")

	ping.GET("/", func(c *gin.Context) {
		keyword := c.Query("keyword")
		result := services.SearchService.FindCardBenefits(keyword)
		c.JSON(http.StatusOK, result)
	})
}
