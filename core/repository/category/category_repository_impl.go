package category

import (
	"context"
	"github.com/debbysa/moleo/core/entity"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) GetCategoryList(ctx context.Context) ([]*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryRepository) GetCategoryById(ctx context.Context, id string) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryRepository) CreateCategory(ctx context.Context, category entity.Category) (*entity.Category, error) {

	result := c.db.Table("categories").Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (c *categoryRepository) UpdateCategory(ctx context.Context, categoryId int, category entity.Category) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryRepository) DeleteCategory(ctx context.Context, categoryId int) ([]*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}
