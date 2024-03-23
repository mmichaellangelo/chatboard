package api

import (
	"chatboard/db"
	"fmt"
	"net/http"
	"regexp"
)

type MessageHandler struct {
	db *db.DBPool
}

var (
	MessageRE = regexp.MustCompile(`^\/messages\/$`)
)

func NewMessageHandler(db *db.DBPool) http.Handler {
	return &MessageHandler{db: db}
}

func (h *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	switch {
	case r.Method == http.MethodPost && MessageRE.MatchString(r.URL.Path):
		err := h.NewMessage(r)
		if err != nil {
			fmt.Println("Error creating message:", err)
			return
		}
	}
}

type Message struct {
	Body string `json:"messageBody"`
}

func (h *MessageHandler) NewMessage(r *http.Request) error {
	r.ParseMultipartForm(0)
	fmt.Println(r.FormValue("messageBody"))
	return nil
}
