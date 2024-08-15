package controllers

import (
	"auth-mux-c09/database"
	"auth-mux-c09/models"
	"database/sql"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

const SecretKey = "secreto"

func Register(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//	bcrypt
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	_, err = database.DB.Exec(
		"INSERT INTO  users (name, email, password) VALUES (?, ?, ?)",
		data["name"], data["email"], password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)

}

func Login(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User

	err = database.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", data["email"]).Scan(&user.Id, &user.Password)
	if err == sql.ErrNoRows {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//	Verificamos hash Comparar los password
	//https://bcrypt.online/
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		http.Error(w, "credenciales incorrectas", http.StatusBadRequest)
		return
	}

	//Generamos un JWT
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		http.Error(w, "could not login", http.StatusInternalServerError)
		return
	}

	//	Seteo del JWT en un cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	json.NewEncoder(w).Encode(map[string]string{"message": "success"})
}

func User(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "unauthenticated", http.StatusUnauthorized)
		return
	}

	//	Obtener info del JWT
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		http.Error(w, "unauthenticated", http.StatusUnauthorized)
		return
	}
	claims := token.Claims.(*jwt.StandardClaims)

	//	Consultamos data en la DB mediante el id del User, obtenido del JWT
	var user models.User

	err = database.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", claims.Issuer).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	json.NewEncoder(w).Encode(map[string]string{"message": "success"})
}
