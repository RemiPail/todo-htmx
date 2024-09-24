package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/RemiPail/todo-htmx/handlers"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return
	}

	router := chi.NewMux()

	router.Handle("/*", public())

	router.Get("/", handlers.Make(handlers.HandleHomeIndex))
	router.Get("/todos", handlers.Make(handlers.HandleTodosIndex))
	router.Post("/api/todos", handlers.Make(handlers.HandleCreateTodo))
	router.Delete("/api/todos/{id}", handlers.Make(handlers.HandleDeleteTodo))
	router.Put("/api/todos/{id}", handlers.Make(handlers.HandleUpdateTodo))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("Listening on", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
