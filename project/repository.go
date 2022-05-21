package project

import (
	"cgr/models"
	"context"
)

type Repository interface {
	Create(context.Context, models.Project) error
	Update(context.Context, models.Project) error
	Delete(context.Context, int) error
	Get(context.Context, int) (models.Project, error)
	GetAllforClient(context.Context) (map[int]models.Project, error)
	GetAllforAdmin(context.Context) (map[int]models.Project, error)
}
