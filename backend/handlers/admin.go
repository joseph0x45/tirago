package handlers

import (
	"backend/store"
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
)

type AdminHandler struct {
	store *store.AdminStore
}

func NewAdminHandler(
	store *store.AdminStore,
) *AdminHandler {
	return &AdminHandler{
		store: store,
	}
}

func (h *AdminHandler) Login(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	admin, err := h.store.GetAdminByUsername(payload.Username)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if admin == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !utils.PasswordMatchesHash(payload.Password, admin.Password) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {

}
