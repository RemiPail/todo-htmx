package handlers

import (
	"net/http"

	"github.com/RemiPail/todo-htmx/views/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
