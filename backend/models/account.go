package models

import "time"

type Account struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateCreated time.Time `json:"date_created"`
}

type NewAccountPostDetails struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
