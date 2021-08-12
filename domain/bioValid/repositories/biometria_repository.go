package repositories

import (
	"database/sql"

	"github.com/matheuscamarques/biovalid/domain/bioValid/entity"
)

type IBiometriaRepository interface {
	GetBiometria(id int) (entity.Biometria, error)
	GetBiometrias() ([]entity.Biometria, error)
	SaveBiometria(biometria entity.Biometria) (int64, error)
	DeleteBiometria(id int) error
}

// BiometriaRepository
// CREATE TABLE "biometria" (
//	"id" int PRIMARY KEY AUTOINCREMENT NOT NULL,
//	"rg" varchar,
//	"cpf" int
//  );
type BiometriaRepository struct {
	IBiometriaRepository
	conn *sql.DB
}

func (repository BiometriaRepository) GetBiometria(id int) (entity.Biometria, error) {
	biometria := entity.Biometria{}
	err := repository.conn.QueryRow(`SELECT * FROM biometria WHERE id = $1`, id).Scan(&biometria)
	return biometria, err
}

func (repository BiometriaRepository) GetBiometrias() ([]entity.Biometria, error) {
	rows, err := repository.conn.Query(`SELECT * FROM biometria`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	biometrias := []entity.Biometria{}
	for rows.Next() {
		biometria := entity.Biometria{}
		err := rows.Scan(&biometria.ID, &biometria.Rg, &biometria.Cpf)
		if err != nil {
			return nil, err
		}
		biometrias = append(biometrias, biometria)
	}
	return biometrias, nil
}

func (repository BiometriaRepository) SaveBiometria(biometria entity.Biometria) (int64, error) {
	result, err := repository.conn.Exec(`INSERT INTO biometria (rg, cpf) VALUES ($1, $2)`, biometria.Rg, biometria.Cpf)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (repository BiometriaRepository) DeleteBiometria(id int) error {
	_, err := repository.conn.Exec(`DELETE FROM biometria WHERE id = $1`, id)
	return err
}
