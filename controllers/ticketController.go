package controllers

import (
	"airline/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetFreeSeatsOnTicket(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "flightId"))

	seats, err := models.GetFreeSeats(id)
	if err != nil {
		result, _ := json.Marshal(map[string]interface{}{
			"message": err,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	result, _ := json.Marshal(map[string]interface{}{
		"message": "ok",
		"seats":   seats,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	return
}

func CreateNewTicket(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	inputUserId, _ := strconv.Atoi(data["userId"])
	inputPrice, _ := strconv.ParseFloat(data["price"], 32)
	inputFlightId, _ := strconv.Atoi(data["flightId"])
	inputSeatCode := data["seatCode"]

	err := models.CreateTicket(inputFlightId, inputUserId, float32(inputPrice), inputSeatCode)
	if err != nil {
		fmt.Println("ticketCon", err)
		result, _ := json.Marshal(map[string]interface{}{
			"message": err,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(result)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetBoughtTickets(w http.ResponseWriter, r *http.Request) {
	inputId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tickets, err := models.GetAllPersonalTickets(inputId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, _ := json.Marshal(map[string]interface{}{
		"tickets": tickets,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
