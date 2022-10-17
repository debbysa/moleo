package api

import "net/http"

func buildGetCategoryByIdError(w http.ResponseWriter, err error) {
	switch err {
	default:
		buildInternalServerResponse(w, err)
	}
}
