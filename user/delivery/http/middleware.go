package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("session")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		access, err := h.usecase.GetByToken(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if !access {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
