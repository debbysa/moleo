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
	categoryList, err := c.categoryRepository.GetCategoryList(ctx)
	if err != nil {
		return nil, err
	}
	return categoryList, nil
}

func (c *categoryUsecase) GetCategoryById(ctx context.Context, id int) (*entity.Category, error) {
	categoryData, err := c.categoryRepository.GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}
	return categoryData, nil
}

func (c *categoryUsecase) CreateCategory(ctx context.Context, category entity.CategoryRequest) (*entity.CategoryRequest, error) {
	createCategory, err := c.categoryRepository.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return createCategory, nil
}

func (c *categoryUsecase) UpdateCategory(ctx context.Context, categoryId int, category entity.CategoryRequest) (*entity.CategoryRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryUsecase) DeleteCategory(ctx context.Context, categoryId int) ([]*entity.CategoryRequest, error) {
	//TODO implement me
	panic("implement me")
}
