package module

import (
	"context"
	"github.com/debbysa/moleo/core/entity"
	"github.com/debbysa/moleo/core/repository/category"
)

type categoryUsecase struct {
	categoryRepository category.CategoryRepository
}

func NewCategoryUsecase(categoryRepository category.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{categoryRepository}
}

func (c *categoryUsecase) GetCategoryList(ctx context.Context) ([]*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryUsecase) GetCategoryById(ctx context.Context, id string) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryUsecase) CreateCategory(ctx context.Context, category entity.Category) (*entity.Category, error) {
	createCategory, err := c.categoryRepository.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return createCategory, nil
}

func (c *categoryUsecase) UpdateCategory(ctx context.Context, categoryId int, category entity.Category) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryUsecase) DeleteCategory(ctx context.Context, categoryId int) ([]*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}
