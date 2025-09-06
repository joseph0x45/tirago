package store

import (
	"backend/models"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SessionStore struct {
	db *sqlx.DB
}

func NewSessionStore(db *sqlx.DB) *SessionStore {
	return &SessionStore{db: db}
}

func (s *SessionStore) InsertSession(session *models.Session) error {
	const query = `
    insert into sessions (
      id, session_type, user_id, valid
    )
    values (
      :id, :session_type, :user_id, :valid
    );
  `
	_, err := s.db.NamedExec(query, session)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to insert new session: %w", err)
	}
	return nil
}

func (s *SessionStore) GetSessionByID(sessionID string) (*models.Session, error) {
	const query = `
    select * from sessions where id=$1
  `
	dbSession := &models.Session{}
	err := s.db.Get(dbSession, query, sessionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[ERROR] Failed to get session by ID: %w", err)
	}
	return dbSession, nil
}

func (s *SessionStore) InvalidateSession(sessionID string) error {
	const query = `
    update sessions set valid=false where id=$1
  `
	_, err := s.db.Exec(query, sessionID)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to invalid session: %w", err)
	}
	return nil
}
