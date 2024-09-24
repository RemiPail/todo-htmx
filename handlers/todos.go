package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/RemiPail/todo-htmx/models"
	"github.com/RemiPail/todo-htmx/views/todos"
	"github.com/go-chi/chi"
)

var todosList []models.Todo = []models.Todo{
	{ID: 1, Title: "todo 1", Done: false},
	{ID: 2, Title: "todo 2", Done: true},
}

func HandleTodosIndex(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, todos.Index(todosList))
}

func HandleCreateTodo(w http.ResponseWriter, r *http.Request) error {
	newTodo := models.Todo{
		ID:    len(todosList) + 1,
		Title: r.FormValue("title"),
		Done:  false,
	}
	todosList = append(todosList, newTodo)

	// Sleep to simulate a long task
	time.Sleep(time.Second * 2)

	return Render(w, r, todos.TodoComponent(newTodo))
}

func HandleDeleteTodo(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}
	for i, todo := range todosList {
		if todo.ID == id {
			todosList = append(todosList[:i], todosList[i+1:]...)
			break
		}
	}
	return nil
}

func HandleUpdateTodo(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	for i, todo := range todosList {
		if todo.ID == id {
			todosList[i].Done = !todosList[i].Done
			return Render(w, r, todos.TodoComponent(todosList[i]))
		}
	}

	return errors.New("todo not found")
}
