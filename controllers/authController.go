package controllers

import (
	"airline/models"
	"airline/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		panic(err)
	}

	inputName := data["name"]
	inputEmail := data["email"]
	inputPassword := data["password"]

	if err := models.Validate(inputEmail, inputPassword); err != nil {
		result, _ := json.Marshal(map[string]interface{}{
			"message": err,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	user, err := models.FindUserByEmail(inputEmail)
	if user.Email != "" {
		result, _ := json.Marshal(map[string]string{
			"message": "already exist",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	user = models.CreateUser(inputName, inputPassword, inputEmail)
	err = user.AddUser()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't add the user to db"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("registration has completed"))
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	inputEmail := data["email"]
	inputPassword := data["password"]

	user, err := models.FindUserByEmail(inputEmail)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user not found"))
		return
	}

	if err := user.ValidatePassword(inputPassword); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect password"))
		return
	}

	token, err := utils.GenereteToken(user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("can't login"))
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour),
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	result, _ := json.Marshal(map[string]interface{}{
		"message": "success login",
		"id":      user.Id,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		Path:     "/",
		HttpOnly: false,
	}

	http.SetCookie(w, &cookie)

	result, _ := json.Marshal(map[string]string{
		"message": "success",
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
