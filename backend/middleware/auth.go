package middleware

import (
	"backend/store"
	"context"
	"log"
	"net/http"
)

type AuthMiddleware struct {
	adminStore   *store.AdminStore
	sessionStore *store.SessionStore
}

func NewAuthmiddleware(
	adminStore *store.AdminStore,
	sessionStore *store.SessionStore,
) *AuthMiddleware {
	return &AuthMiddleware{
		adminStore:   adminStore,
		sessionStore: sessionStore,
	}
}

func (m *AuthMiddleware) SessionAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.Header.Get("session")
		if sessionID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		session, err := m.sessionStore.GetSessionByID(sessionID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !session.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "session", session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
