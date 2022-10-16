package category

import (
	"context"
	"github.com/debbysa/moleo/core/entity"
)

type CategoryRepository interface {
	GetCategoryList(ctx context.Context) ([]*entity.Category, error)
	GetCategoryById(ctx context.Context, id int) (*entity.Category, error)
	CreateCategory(ctx context.Context, category entity.CategoryRequest) (*entity.CategoryRequest, error)
	UpdateCategory(ctx context.Context, categoryId int, category entity.CategoryRequest) (*entity.CategoryRequest, error)
	DeleteCategory(ctx context.Context, categoryId int) ([]*entity.CategoryRequest, error)
}
