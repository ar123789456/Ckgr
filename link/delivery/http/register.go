package http

import (
	"cgr/link"
	"cgr/tool/logger"

	"github.com/gin-gonic/gin"
)

func RegisterLink(router *gin.Engine, uc link.UseCase, logg *logger.Logger) {
	h := NewHandler(uc, logg)
	linkEndpoints := router.Group("/link")
	{
		linkEndpoints.POST("/", h.Create)
		linkEndpoints.GET("/", h.GetAll)
		linkEndpoints.DELETE("/", h.Delete)
	}
}
