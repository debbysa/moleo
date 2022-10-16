package api

import (
	"github.com/debbysa/moleo/core/entity"
	"github.com/debbysa/moleo/core/module"
	"net/http"
)

type categoryHandler struct {
	categoryUsecase module.CategoryUsecase
}
type CategoryHandler interface {
	GetCategoryList(w http.ResponseWriter, r *http.Request)
	GetCategoryById(w http.ResponseWriter, r *http.Request)
	CreateCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
}

func NewCategoryHandler(categoryUsecase module.CategoryUsecase) CategoryHandler {
	return &categoryHandler{categoryUsecase}
}

func (c *categoryHandler) GetCategoryList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	categoryCreateRequest := entity.Category{}

	ReadFromRequestBody(w, r, &categoryCreateRequest)

	categoryResponse, err := c.categoryUsecase.CreateCategory(r.Context(), categoryCreateRequest)

	if err != nil {
		switch err {
		default:
			buildGetCategoryByIdError(w, err)
			return
		}
	}
	buildSuccessResponse(w, &categoryResponse)
}

func (c *categoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
