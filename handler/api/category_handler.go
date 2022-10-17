package api

import (
	"github.com/debbysa/moleo/core/entity"
	"github.com/debbysa/moleo/core/module"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	categoryList, err := c.categoryUsecase.GetCategoryList(r.Context())
	if err != nil {
		buildInternalServerResponse(w, err)
		return
	}
	buildSuccessResponse(w, categoryList)
}

func (c *categoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]
	parseCategoryID, err := strconv.Atoi(categoryID)
	if err != nil {
		buildInternalServerResponse(w, err)
		return
	}

	categoryData, err := c.categoryUsecase.GetCategoryById(r.Context(), parseCategoryID)
	if err != nil {
		switch err {
		default:
			buildGetCategoryByIdError(w, err)
			return
		}
	}
	buildSuccessResponse(w, categoryData)
}

func (c *categoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	categoryCreateRequest := entity.CategoryRequest{}

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
	vars := mux.Vars(r)
	categoryId := vars["id"]
	parseCategoryId, err := strconv.Atoi(categoryId)
	if err != nil {
		buildInternalServerResponse(w, err)
		return
	}

	categoryUpdateRequest := entity.CategoryRequest{}
	ReadFromRequestBody(w, r, &categoryUpdateRequest)
	categoryResponse, err := c.categoryUsecase.UpdateCategory(r.Context(), parseCategoryId, categoryUpdateRequest)
	if err != nil {
		switch err {
		default:
			buildGetCategoryByIdError(w, err)
			return
		}
	}
	buildSuccessResponse(w, &categoryResponse)

}

func (c *categoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["id"]
	parseCategoryId, err := strconv.Atoi(categoryId)
	if err != nil {
		buildInternalServerResponse(w, err)
		return
	}
	err = c.categoryUsecase.DeleteCategory(r.Context(), parseCategoryId)
	if err != nil {
		buildInternalServerResponse(w, err)
	}
	buildSuccessResponse(w, nil)
}
