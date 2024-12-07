package entities

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	ID           pgtype.UUID
	Email        string
	Password     string
	SessionToken string
}
