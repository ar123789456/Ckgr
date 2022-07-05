package http

import (
	"cgr/models"
	"cgr/project"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase project.UseCase
}

func NewHandler(usecase project.UseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type createInput struct {
	ID     int     `json:"id"`
	Base   frame   `json:"base"`
	Enable bool    `json:"enable"`
	All    []frame `json:"all"`
}
type frame struct {
	Title   map[models.Language]string `json:"title"`
	Image   string                     `json:"image"`
	Content map[models.Language]string `json:"content"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := h.usecase.Post(c.Request.Context(), bindModel(*inp))

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}

type getInput struct {
	ID int `json:"id"`
}

func (h *Handler) Get(c *gin.Context) {
	inp := new(getInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	model, err := h.usecase.Get(c.Request.Context(), inp.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	jnews := bindOutputProject(model)
	c.JSON(http.StatusOK, &jnews)
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(getInput)
	if err := c.BindJSON(inp); err != nil {
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
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := h.usecase.Update(c.Request.Context(), bindModel(*inp))

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetAllForClient(c *gin.Context) {
	models, err := h.usecase.GetAllforClient(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	jnews := bindProjects(models)
	c.JSON(http.StatusOK, &jnews)
}

func (h *Handler) GetAllForAdmin(c *gin.Context) {
	models, err := h.usecase.GetAllforAdmin(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	jnews := bindProjects(models)
	c.JSON(http.StatusOK, &jnews)
}

type output struct {
	ID     int     `json:"id"`
	Base   frame   `json:"base"`
	Enable bool    `json:"enable"`
	All    []frame `json:"all"`
}

func bindProjects(pr []models.Project) []output {
	out := []output{}
	for _, v := range pr {
		out = append(out, bindOutputProject(v))
	}
	return out
}

func bindOutputProject(proj models.Project) output {
	all := []frame{}
	for _, v := range proj.All {
		all = append(all, bindOutPutFrame(v))
	}

	return output{
		ID:     proj.ID,
		Base:   bindOutPutFrame(proj.Base),
		Enable: proj.Enable,
		All:    all,
	}
}

func bindOutPutFrame(f models.Frame) frame {
	return frame{
		Title:   f.Title,
		Image:   f.Image,
		Content: f.Content,
	}
}

func bindModel(ci createInput) models.Project {
	all := []models.Frame{}
	for _, v := range ci.All {
		all = append(all, bindFrame(v))
	}

	return models.Project{
		Base:   bindFrame(ci.Base),
		Enable: ci.Enable,
		All:    all,
	}
}
func bindFrame(f frame) models.Frame {
	return models.Frame{
		Title:   f.Title,
		Image:   f.Image,
		Content: f.Content,
	}
}
