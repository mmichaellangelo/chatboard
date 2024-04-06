package api

import (
	"chatboard/db"
	"chatboard/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

type AccountHandler struct {
	db *db.DBPool
}

var (
	AccountRE             = regexp.MustCompile(`^\/accounts\/$`)
	AccountREWithID       = regexp.MustCompile(`^\/accounts\/id\/[0-9]+$`)
	AccountREWithUsername = regexp.MustCompile(`^\/accounts\/username\/[A-z0-9-_]+$`)
)

func NewAccountHandler(db *db.DBPool) http.Handler {
	return &AccountHandler{db: db}
}

func (h *AccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	// GET All Accounts
	case r.Method == http.MethodGet && AccountRE.MatchString(r.URL.Path):
		fmt.Println("GET ALL ACC ROUTE")
		accounts, err := h.GetAllAccounts()
		accountsJSON, err := json.Marshal(accounts)
		if err != nil {
			fmt.Println("Error marshalling accounts:", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(accountsJSON)

	// GET Account by ID
	case r.Method == http.MethodGet && AccountREWithID.MatchString(r.URL.Path):
		idString := r.URL.Path[len("/accounts/id/"):]
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			fmt.Println("Error parsing id:", err)
			return
		}
		account, err := h.GetAccountByID(id)
		if err != nil {
			fmt.Println("Error querying account by ID:", err)
			return
		}
		if account.Username == "" {
			w.WriteHeader(404)
			return
		}
		accountJSON, err := json.Marshal(account)
		w.Write(accountJSON)

	// CREATE Account
	case r.Method == http.MethodPost && AccountRE.MatchString(r.URL.Path):
		r.ParseMultipartForm(0)
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if username == "" || email == "" || password == "" {
			w.WriteHeader(400)
			return
		}

		err := h.CreateAccount(username, email, password)
		if err != nil {
			fmt.Println("Error creating new account:", err)
			w.WriteHeader(500)
			return
		}

	default:
		return
	}
}

func (h *AccountHandler) GetAllAccounts() ([]models.Account, error) {
	var accounts []models.Account
	rows, err := h.db.Pool.Query(context.Background(), "SELECT id, username, email, password, date_created FROM account")
	if err != nil {
		fmt.Println("Error querying accounts:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account.ID, &account.Username, &account.Email, &account.Password, &account.DateCreated)
		if err != nil {
			fmt.Println("Error querying row:", err)
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (h *AccountHandler) GetAccountByID(id int64) (models.Account, error) {
	fmt.Println("GETTING ACCOUNT #", id)
	var account models.Account
	rows, err := h.db.Pool.Query(context.Background(), `SELECT id, username, email, password, date_created FROM account WHERE id=$1`, id)
	if err != nil {
		fmt.Println("Error querying account:", err)
		return account, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&account.ID, &account.Username, &account.Email, &account.Password, &account.DateCreated)
	}
	if err != nil {
		fmt.Println("Error querying row:", err)
		return account, err
	}

	return account, nil
}

var (
	UsernameRE = regexp.MustCompile(`^[A-z0-9-]+$`)
	EmailRE    = regexp.MustCompile(`^[A-z0-9-]+@[A-z0-9-]+\.[A-z]{2,16}$`)
	PasswordRE = regexp.MustCompile(`^.{8,64}$`)
)

func (h *AccountHandler) CreateAccount(username string, email string, password string) error {
	if !(UsernameRE.MatchString(username) && EmailRE.MatchString(email) && PasswordRE.MatchString(password)) {
		return fmt.Errorf("Error: fields do not meet requirements")
	}
	rows, err := h.db.Pool.Query(context.Background(), "INSERT INTO account (username, email, password, date_created) VALUES ($1, $2, $3, NOW())", username, email, password)
	if err != nil {
		return fmt.Errorf("Error inserting into database: %v", err)
	}
	fmt.Println(rows)
	return nil
}
