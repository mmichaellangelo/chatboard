package api

import (
	"chatboard/auth"
	"chatboard/db"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type LoginHandler struct {
	db *db.DBPool
}

var (
	LoginRE = regexp.MustCompile(`^\/login\/$`)
)

func NewLoginHandler(db *db.DBPool) http.Handler {
	return &LoginHandler{db: db}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && LoginRE.MatchString(r.URL.Path):
		token, err := h.HandleLogin(r)
		if err != nil {
			fmt.Println("Error logging in:", err)
			w.WriteHeader(500)
			return
		}
		if token == "" {
			fmt.Println("Error: blank token")
			w.WriteHeader(500)
			return
		}

		responseData := map[string]interface{}{
			"token": token,
		}

		jsonData, err := json.Marshal(responseData)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func (h *LoginHandler) HandleLogin(r *http.Request) (string, error) {
	r.ParseMultipartForm(0)
	username := r.FormValue("username")
	password := r.FormValue("password")
	token, err := h.AuthorizeUser(username, password)
	if err != nil {
		return "", err
	}
	if token == "" {
		fmt.Println("Auth failed.")
		return "", nil
	}
	fmt.Println("Auth success! Token:", token)
	return token, nil
}

type AuthToken struct {
	token string
}

func (h *LoginHandler) AuthorizeUser(username string, password string) (string, error) {
	rows, err := h.db.Pool.Query(context.Background(), "SELECT password FROM account WHERE username=$1", username)
	if err != nil {
		return "", err
	}
	var passwordFromDB string
	if rows.Next() {
		err := rows.Scan(&passwordFromDB)
		if err != nil {
			return "", err
		}
	}
	if password == passwordFromDB {
		fmt.Println("Auth Success!")
		token, err := auth.CreateToken(username)
		if err != nil {
			return "", fmt.Errorf("Error creating token: %v", err)
		}

		return token, nil
	} else {
		fmt.Println("Auth Denied.")
		return "", nil
	}
}
