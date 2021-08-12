package repositories

import (
	"database/sql"
)

type AccessTokenRepository struct {
	conn *sql.DB
}
