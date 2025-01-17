package routes

import (
	"encoding/json"
	"log"
	"net/http"
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
