package repositories

import (
	"database/sql"

	"github.com/matheuscamarques/biovalid/domain/bioValid/entity"
)

type AccessTokenRepository struct {
	conn *sql.DB
}

func (r AccessTokenRepository) Create(accessToken *entity.AccessToken) error {
	_, err := r.conn.Exec(`INSERT INTO access_tokens (id_user_api,token,expired_at) VALUES ($1,$2,$3)`, accessToken.IDUserApi, accessToken.Expired_at)
	return err
}

func (r AccessTokenRepository) GetUserApiByToken(token string) (*entity.UserApi, error) {
	var accessToken entity.AccessToken
	err := r.conn.QueryRow(`SELECT * FROM access_tokens WHERE token = $1`, token).Scan(&accessToken)

	if err != nil {
		return nil, err
	}
	var userApi entity.UserApi
	err = r.conn.QueryRow(`SELECT * FROM user_api WHERE id = $1`, accessToken.IDUserApi).Scan(&userApi)

	if err != nil {
		return nil, err
	}
	return &userApi, nil
}

func (r AccessTokenRepository) GetAll() ([]entity.AccessToken, error) {
	var accessTokens []entity.AccessToken
	rows, err := r.conn.Query(`SELECT * FROM access_tokens`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var accessToken entity.AccessToken
		err := rows.Scan(&accessToken.IDUserApi, &accessToken.Token, &accessToken.Expired_at)
		if err != nil {
			return nil, err
		}
		accessTokens = append(accessTokens, accessToken)
	}

	return accessTokens, nil
}

func (r AccessTokenRepository) Delete(id int) error {
	_, err := r.conn.Exec(`DELETE FROM access_tokens WHERE id = $1`, id)
	return err
}
