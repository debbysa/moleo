package module

import (
	"context"
	"github.com/debbysa/moleo/core/entity"
)

type CategoryUsecase interface {
	GetCategoryList(ctx context.Context) ([]*entity.Category, error)
	GetCategoryById(ctx context.Context, id string) (*entity.Category, error)
	CreateCategory(ctx context.Context, category entity.Category) (*entity.Category, error)
	UpdateCategory(ctx context.Context, categoryId int, category entity.Category) (*entity.Category, error)
	DeleteCategory(ctx context.Context, categoryId int) ([]*entity.Category, error)
}
