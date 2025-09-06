package store

import (
	"backend/models"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AdminStore struct {
	db *sqlx.DB
}

func NewAdminStore(
	db *sqlx.DB,
) *AdminStore {
	return &AdminStore{db: db}
}

func (s *AdminStore) InsertAdmin(admin *models.Admin) error {
	const query = `
    insert into admins(
      id, username, password
    )
    values (
      :id, :username, :password
    )
  `
	_, err := s.db.NamedExec(query, admin)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to insert new admin account: %w", err)
	}
	return nil
}

func (s *AdminStore) GetAdminByID(id string) (*models.Admin, error) {
	const query = `
    select * from admins where id=$1
  `
	admin := &models.Admin{}
	err := s.db.Get(admin, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[ERROR] Failed to get admin by id: %w", err)
	}
	return admin, nil
}

func (s *AdminStore) GetAdminByUsername(username string) (*models.Admin, error) {
	const query = `
    select * from admins where username=$1
  `
	admin := &models.Admin{}
	err := s.db.Get(admin, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[ERROR] Failed to get admin by username: %w", err)
	}
	return admin, nil
}

func (s *AdminStore) CountAdminsInDB() (*int, error) {
	var count = 0
	err := s.db.QueryRow("select count(*) from admins").Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to count admin accounts: %w", err)
	}
	return &count, nil
}

func (s *AdminStore) EnsureAdminAccountExists() error {
	adminCount, err := s.CountAdminsInDB()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to ensure admin account exists: %w", err)
	}
	if *adminCount != 0 {
		return nil
	}
	//create new admin
	newAdmin := &models.Admin{
		ID:       uuid.NewString(),
		Username: "admin",
		Password: os.Getenv("DEFAULT_ADMIN_PASSWORD"),
	}
	err = s.InsertAdmin(newAdmin)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to insert new admin account: %w", err)
	}
	return nil
}

func (s *AdminStore) GetAllAdmins() ([]models.Admin, error) {
	admins := make([]models.Admin, 0)
	const query = "select * from admins"
	err := s.db.Select(&admins, query)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to get all admin accounts: %w", err)
	}
	return admins, nil
}

func (s *AdminStore) UpdateAdminPassword(password, id string) error {
	const query = `
    update admins set password=$1 where id=$2
  `
	_, err := s.db.Exec(query, password, id)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to change admin password: %w", err)
	}
	return nil
}
