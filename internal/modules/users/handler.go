package users

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
    Repo *Repository
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux){
    mux.HandleFunc("GET /users", h.Get)
    mux.HandleFunc("POST /users", h.Create)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
    users, err := h.Repo.GetAll()
    if err != nil {
        http.Error(w, "failed to retrieve users", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(users); err != nil {
        return
    }
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request){
    defer r.Body.Close()
    var user User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "name, email is required", http.StatusBadRequest)
        return
    }

    if user.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

    if err := h.Repo.Store(&user); err != nil{
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(http.StatusCreated)
    _ = json.NewEncoder(w).Encode(user)
}