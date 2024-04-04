package main

import (
	"chatboard/api"
	"chatboard/db"
	"context"
	"fmt"
	"net/http"
)

func main() {

	// Establish DB Connection Pool

	pool, err := db.Connect(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	// Establish API routes and start server

	err = createRoutesAndServe(pool)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func createRoutesAndServe(db *db.DBPool) error {
	mux := http.NewServeMux()
	mux.Handle("/", api.NewHomeHandler(db))
	mux.Handle("/messages/", api.NewMessageHandler(db))
	mux.Handle("/accounts/", api.NewAccountHandler(db))
	mux.Handle("/login/", api.NewLoginHandler(db))

	err := http.ListenAndServe(":8080", mux)
	return err
}
