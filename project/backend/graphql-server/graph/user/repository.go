package user

import (
	"context"
	"errors"
	"fmt"
	"mindmap-backend/graphql-server/graph/structures"
	"mindmap-backend/graphql-server/graph/utils"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	userTable = "users"
)

var (
	insertUserColumns = []string{"id", "username", "password"}
)

type DBRepositoryUser struct {
	db *sqlx.DB
}

func NewDBRepositoryUser(db *sqlx.DB) *DBRepositoryUser {
	return &DBRepositoryUser{
		db: db,
	}
}

func (r *DBRepositoryUser) CreateUser(ctx context.Context, userEntity structures.UserEntity) error {
	log := ctx.Value(utils.Logger).(*logrus.Entry)

	tx, err := r.db.Beginx()
	if err != nil {
		log.Error(err)
		return err
	}
	defer tx.Rollback()

	stmt := fmt.Sprintf(`INSERT INTO %s(%s) VALUES (?, ?, ?)`, userTable, strings.Join(insertUserColumns, ", "))
	query := sqlx.Rebind(sqlx.DOLLAR, stmt)
	_, err = tx.Exec(query, userEntity.Id, userEntity.Username, userEntity.Password)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			message := "error creating user (%s) because of: %s"
			switch pgErr.Code {
			case utils.UNIQUE_VIOLATION:
				message = fmt.Sprintf(message, userEntity.Username, "user with this username already exists")
				log.Warn(message)
				return errors.New(message)
			default:
				message = fmt.Sprintf(message, userEntity.Username,
					fmt.Sprintf("database error [%s]", pgErr.Code))
				log.Warn(message)
				return errors.New(message)
			}
		}

		return errors.New(fmt.Sprintf(utils.UNCLASIFIED_ERROR, err))
	}

	err = tx.Commit()
	if err != nil {
		log.Error(err)
	}

	return err
}
