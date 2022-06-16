package project

import (
	"cgr/models"
	"context"
)

type UseCase interface {
	Post(context.Context, models.Project) error
	Create(context.Context, models.Project) error
	Update(context.Context, models.Project) error
	Delete(context.Context, string) error
	Get(context.Context, string) (models.Project, error)
	GetAllforClient(context.Context) ([]models.Project, error)
	GetAllforAdmin(context.Context) ([]models.Project, error)
}
