package utils

import (
	"data"
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
}
