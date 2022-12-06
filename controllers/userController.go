package controllers

import (
	"airline/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	idFromToken := fmt.Sprintf("%v", r.Context().Value("idFromToken"))
	userId, _ := strconv.Atoi(idFromToken)

	if id != userId {
		result, _ := json.Marshal(map[string]string{
			"message": "access denied",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	user, _ := models.FindUserById(uint(id))
	result, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	inputEmail := data["email"]
	inputName := data["name"]
	inputLastName := data["lastname"]
	inputPatronymic := data["patronymic"]
	inputDocument := data["document"]
	inputPhone := data["phone"]

	user, err := models.FindUserByEmail(inputEmail)
	if err != nil {
		result, _ := json.Marshal(map[string]interface{}{
			"message": err,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	err = user.UpdateInfo(inputName, inputLastName, inputPatronymic, inputDocument, inputPhone)
	if err != nil {
		result, _ := json.Marshal(map[string]interface{}{
			"message": err,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(result)
		return
	}

	result, _ := json.Marshal(map[string]string{
		"message": "ok",
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	inputId, _ := strconv.Atoi(data["id"])

	err := models.DeleteUserById(inputId)
	if err != nil {
		result, _ := json.Marshal(map[string]interface{}{
			"message": err,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(result)
		return
	}

	result, _ := json.Marshal(map[string]string{
		"message": "ok",
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
