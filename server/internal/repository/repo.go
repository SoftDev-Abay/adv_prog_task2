package repository

import (
	"database/sql"
	"renting/models"
)

type Store interface {
	Connection() *sql.DB
	GetBuildings() ([]models.Building, error)
	GetCountBuildings(...string) (int, error)
	GetBuildingsInRange(int, int, ...string) ([]models.Building, error)
	Register(models.User) error
	LoginUser(username, password string) (models.User, error)
	InsertBuilding(building models.Building) (int, error)
	UserExistsWithEmail(email string) (bool, error)
}
