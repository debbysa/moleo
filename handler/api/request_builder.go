package api

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(w http.ResponseWriter, r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(buildResponse(nil, http.StatusText(http.StatusBadRequest), err))
	}
}
