package http

import (
	"cgr/link"
	"cgr/models"
	"cgr/tool/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	//news usecase
	usecase link.UseCase
	//custom logger
	log *logger.Logger
}

// Create new link handler
func NewHandler(usecase link.UseCase, log *logger.Logger) *Handler {
	return &Handler{
		usecase: usecase,
		log:     log,
	}
}

type input struct {
	Logo  string `json:"logo"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(input)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := h.usecase.Create(c.Request.Context(), toLink(inp))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

type delInp struct {
	ID string `json:"id"`
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(delInp)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := h.usecase.Delete(c.Request.Context(), inp.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) GetAll(c *gin.Context) {
	all, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, toLinkOut(all))
}

func toLink(inp *input) models.Link {
	return models.Link{
		Logo:  inp.Logo,
		Title: inp.Title,
		Url:   inp.Url,
	}
}

func toLinkOut(all []*models.Link) []input {
	ret := []input{}
	for _, v := range all {
		tem := input{
			Logo:  v.Logo,
			Title: v.Title,
			Url:   v.Url,
		}
		ret = append(ret, tem)
	}
	return ret
}
