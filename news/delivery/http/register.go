package http

import (
	"cgr/news"
	"cgr/tool/logger"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc news.UseCase, logger *logger.Logger) {
	h := NewHandler(uc, logger)

	newsRoute := router.Group("/bookmarks")
	{
		newsRoute.POST("", h.Create)
		newsRoute.GET("/single", h.Get)
		newsRoute.DELETE("/single", h.Delete)
		newsRoute.PUT("/single", h.Update)
		newsRoute.GET("/news/admin", h.GetAllForAdmin)
		newsRoute.GET("/news", h.GetAllForClient)
	}
}
