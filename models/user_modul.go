package models

type User struct {
	ID           string `json:"id" db:"id"`
	FullName     string `json:"full_name" db:"full_name"`
	Username     string `json:"username" db:"username"`
	Email        string `json:"email" db:"email"`
	Bio          string `json:"bio" db:"bio"`
	UserType     string `json:"user_type" db:"user_type"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}
