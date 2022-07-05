package handler

import (
	"cgr/models"
	"cgr/user"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var domain = ""

type Handler struct {
	usecase user.UseCase
}

func NewHandler(usecase user.UseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type logInInput struct {
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

func (h *Handler) LogIn(c *gin.Context) {
	login := new(logInInput)
	if err := c.BindJSON(login); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid json"))
		return
	}
	uuid, err := h.usecase.LogIn(c.Request.Context(), models.User{
		Nick:     login.Nick,
		Password: login.Password,
	})
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.SetCookie("session", uuid, 4000, "/", domain, false, false)
	c.Status(http.StatusOK)
}

type createInput struct {
	Nick     string                     `json:"nick"`
	FullName map[models.Language]string `json:"name"`
	Special  map[models.Language]string `json:"special"`
	Password string                     `json:"password"`
}

func (h *Handler) Create(c *gin.Context) {
	user := new(createInput)
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := h.usecase.Create(c.Request.Context(), models.User{
		Nick:     user.Nick,
		FullName: user.FullName,
		Password: user.Password,
		Special:  user.Special,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}
func (h *Handler) Delete(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.usecase.Delete(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

type updateInput struct {
	ID       int                        `json:"id"`
	FullName map[models.Language]string `json:"name"`
	Special  map[models.Language]string `json:"special"`
	Password string                     `json:"password"`
}

func (h *Handler) Update(c *gin.Context) {
	user := new(updateInput)
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := h.usecase.Update(c.Request.Context(), models.User{
		ID:       user.ID,
		FullName: user.FullName,
		Password: user.Password,
		Special:  user.Special,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}
