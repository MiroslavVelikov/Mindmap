package structures

import "github.com/google/uuid"

type UserEntity struct {
	Id       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Password string    `db:"password"`
}
