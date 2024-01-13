package repository

import (
	"database/sql"
	"renting/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetBuildings() ([]models.Building, error)
	Register(models.User) error
	LoginUser(username, password string) (models.User, error)
	InsertBuilding(building models.Building) (int, error)
}
