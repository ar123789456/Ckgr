package http

import (
	"cgr/models"
	"cgr/news"
	"cgr/tool/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Http handler
type Handler struct {
	//news usecase
	usecase news.UseCase
	//custom logger
	log *logger.Logger
}

// Create new news handler
func NewHandler(usecase news.UseCase, log *logger.Logger) *Handler {
	return &Handler{
		usecase: usecase,
		log:     log,
	}
}

type createInput struct {
	Image_url string   `json:"image"`
	Tags      []string `json:"tags"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Enable    bool     `json:"enable"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		h.log.ErrorLogger.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := h.usecase.Post(c.Request.Context(), models.News{
		Image:   inp.Image_url,
		Tags:    inp.Tags,
		Title:   inp.Title,
		Content: inp.Content,
		Enable:  inp.Enable,
	})

	if err != nil {
		h.log.ErrorLogger.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}

type getInput struct {
	ID string `json:"id"`
}

func (h *Handler) Get(c *gin.Context) {
	inp := new(getInput)
	if err := c.BindJSON(inp); err != nil {
		h.log.ErrorLogger.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	model, err := h.usecase.Get(c.Request.Context(), inp.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	jnews := toNew(model)
	c.JSON(http.StatusOK, &jnews)
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(getInput)
	if err := c.BindJSON(inp); err != nil {
		h.log.ErrorLogger.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := h.usecase.Delete(c.Request.Context(), inp.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (h *Handler) Update(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		h.log.ErrorLogger.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := h.usecase.Update(c.Request.Context(), models.News{
		Image:   inp.Image_url,
		Tags:    inp.Tags,
		Title:   inp.Title,
		Content: inp.Content,
		Enable:  inp.Enable,
	})

	if err != nil {
		h.log.ErrorLogger.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetAllForClient(c *gin.Context) {
	models, err := h.usecase.GetAllForClient(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	jnews := toNews(models)
	c.JSON(http.StatusOK, &jnews)
}

func (h *Handler) GetAllForAdmin(c *gin.Context) {
	models, err := h.usecase.GetAllForAdmin(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	jnews := toNews(models)
	c.JSON(http.StatusOK, &jnews)
}

type output struct {
	Enable  bool
	Image   string   `json:"image"`
	Tags    []string `json:"tags"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
}

func toNew(n models.News) output {
	return output{
		Enable:  n.Enable,
		Image:   n.Image,
		Tags:    n.Tags,
		Title:   n.Title,
		Content: n.Content,
	}
}

func toNews(n []models.News) []output {
	ret := []output{}
	for _, v := range n {
		ret = append(ret, toNew(v))
	}
	return ret
}
