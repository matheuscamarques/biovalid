package repositories

import (
	"database/sql"
)

type AccessTokenRepository struct {
	conn *sql.DB
}

func (repo AccessTokenRepository) GetAccessToken(token string) (string, error) {
	var accessToken string
	entity.AccessToken
	err := repo.conn.QueryRow("SELECT access_token FROM access_token WHERE token = $1", token).Scan(&accessToken)
	return accessToken, err
}
