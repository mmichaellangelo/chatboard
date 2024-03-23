package api

import (
	"chatboard/db"
	"net/http"
	"regexp"
)

type HomeHandler struct {
	db *db.DBPool
}

var (
	HomeRE = regexp.MustCompile(`^\/$`)
)

func NewHomeHandler(db *db.DBPool) http.Handler {
	return &HomeHandler{db: db}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && HomeRE.MatchString(r.URL.Path):
		w.Write([]byte("Hi bitch"))
		return
	default:
		w.WriteHeader(404)
	}
}
