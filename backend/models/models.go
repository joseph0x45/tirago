package models

import "time"

type Admin struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type MembershipRequest struct {
	ID            string    `json:"id" db:"id"`
	Email         string    `json:"email" db:"email"`
	Phone         string    `json:"phone" db:"phone"`
	FullName      string    `json:"full_name" db:"full_name"`
	AccountType   string    `json:"account_type" db:"account_type"`
	Status        string    `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	RefusalReason string    `json:"refusal_reason" db:"refusal_reason"`
}

type MembershipRequestDoc struct {
	ID          string `json:"id" db:"id"`
	RequestID   string `json:"request_id" db:"request_id"`
	DocumentURL string `json:"document_url" db:"document_url"`
}

type User struct {
	ID             string      `json:"id" db:"id"`
	Email          string      `json:"email" db:"email"`
	Phone          string      `json:"phone" db:"phone"`
	Town           string      `json:"town" db:"town"`
	Password       string      `json:"password" db:"password"`
	FullName       string      `json:"full_name" db:"full_name"`
	ProfilePicture string      `json:"profile_picture" db:"profile_picture"`
	AccountType    string      `json:"account_type" db:"account_type"`
	CreatedAt      time.Ticker `json:"created_at" db:"created_at"`
}

type UserDoc struct {
	ID          string `json:"id" db:"id"`
	UserID      string `json:"user_id" db:"user_id"`
	DocumentURL string `json:"document_url" db:"document_url"`
}
