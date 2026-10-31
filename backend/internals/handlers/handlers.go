package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/grant1x/todo/internals/database"
)

type Handler struct {
	db *database.DB
}

type Notes struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"` 
}

func New(db *database.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) HandleFetchNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fetching notes..."))

	rows, err := h.db.Pool.Query("SELECT id, title, content FROM notes")

	if err != nil {
		http.Error(w, "failed to query database", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var notes []Notes

	for rows.Next() {
		var n Notes
		if err := rows.Scan(&n.ID, &n.Title, &n.Content); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}

		notes = append(notes, n)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func (h *Handler) HandleCreateNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a note..."))
}

func (h *Handler) HandleDefault(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}