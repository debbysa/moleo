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
	var categoryDetail []*entity.Category
	result := c.db.Table("categories").Find(&categoryDetail)
	if result.Error != nil {
		return nil, result.Error
	}
	return categoryDetail, nil
}

func (c *categoryRepository) GetCategoryById(ctx context.Context, id int) (*entity.Category, error) {
	var category *entity.Category
	result := c.db.Table("categories").First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (c *categoryRepository) CreateCategory(ctx context.Context, category entity.CategoryRequest) (*entity.CategoryRequest, error) {
	result := c.db.Table("categories").Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (c *categoryRepository) UpdateCategory(ctx context.Context, categoryId int, category entity.CategoryRequest) (*entity.Category, error) {
	categoryData := entity.Category{
		ID:   categoryId,
		Name: category.Name,
	}

	result := c.db.Table("categories").Save(&categoryData)
	if result.Error != nil {
		return nil, result.Error
	}
	return &categoryData, nil
}

func (c *categoryRepository) DeleteCategory(ctx context.Context, categoryId int) error {
	category := entity.Category{
		ID: categoryId,
	}
	result := c.db.Table("categories").Delete(&category)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}
