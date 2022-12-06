package controllers

import (
	"airline/models"
	"encoding/json"
	"net/http"
)

func GetFlights(w http.ResponseWriter, r *http.Request) {
	flights, err := models.GetAllFlights()

	if err != nil {
		result, _ := json.Marshal(map[string]string{
			"message": "internal error",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(result)
		return
	}

	result, _ := json.Marshal(flights)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
