package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/l4rma/todo-go/pkg/db/entity"
)

type Handler struct{}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incomming %s request", r.Method)

	tasks, err := taskService.FindAll()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	response, _ := json.Marshal(tasks)
	w.Write(response)
}

func (h *Handler) FindById(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incomming %s request", r.Method)

	id := r.URL.Query().Get("id")
	task, err := taskService.FindById(id)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	response, _ := json.Marshal(task)
	w.Write(response)
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incomming %s request", r.Method)

	task := &entity.Task{}
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		http.Error(w, "Error: Failed to unmarshal event", http.StatusInternalServerError)
		return
	}

	task, err := taskService.Create(task)
	if err != nil {
		http.Error(w, "Error: Failed to create task", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(task)
	w.Write(response)
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incomming %s request", r.Method)

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Error: Missing id", http.StatusBadRequest)
		return
	}

	task, err := taskService.FindById(id)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	title := r.URL.Query().Get("title")
	if title != "" {
		title = task.Title
	}

	description := r.URL.Query().Get("description")
	if description != "" {
		description = task.Description
	}

	completed := r.URL.Query().Get("completed")
	if completed != "" {
		comp, err := strconv.ParseBool(completed)
		if err != nil {
			http.Error(w, "Invalid value for completed", http.StatusBadRequest)
			return
		}
		task.Completed = comp
	}

	task, err := taskService.Update(task)
	if err != nil {
		http.Error(w, "Error: Failed to update task", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(task)
	w.Write(response)
}
