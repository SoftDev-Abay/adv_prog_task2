package repository

import (
	"database/sql"
	"renting/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetBuildings() ([]models.Building, error)
	Register(username, email, password, phoneNum string) error
	LoginUser(username, password string) (models.User, error)
}
