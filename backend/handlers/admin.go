package handlers

import (
	"backend/consts"
	"backend/models"
	"backend/store"
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type AdminHandler struct {
	adminStore   *store.AdminStore
	sessionStore *store.SessionStore
}

func NewAdminHandler(
	adminStore *store.AdminStore,
	sessionStore *store.SessionStore,
) *AdminHandler {
	return &AdminHandler{
		adminStore:   adminStore,
		sessionStore: sessionStore,
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
	admin, err := h.adminStore.GetAdminByUsername(payload.Username)
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
	newSession := &models.Session{
		ID:          uuid.NewString(),
		SessionType: consts.AdminSessionType,
		UserID:      admin.ID,
		Valid:       true,
	}
	err = h.sessionStore.InsertSession(newSession)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(map[string]map[string]string{
		"data": {
			"session": newSession.ID,
		},
	})
	if err != nil {
		log.Println("[ERROR] Failed to marshall session ID data: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *AdminHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	currentSession, ok := r.Context().Value("session").(*models.Session)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if currentSession.SessionType != consts.AdminSessionType {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	payload := struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("[ERROR] Failed to decode body: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	adminUser, err := h.adminStore.GetAdminByID(currentSession.UserID)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if adminUser == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !utils.PasswordMatchesHash(payload.CurrentPassword, adminUser.Password) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newPasswordHash, err := utils.HashPassword(payload.NewPassword)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.adminStore.UpdateAdminPassword(newPasswordHash, adminUser.ID)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
