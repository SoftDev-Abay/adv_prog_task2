package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"renting/models"

	"golang.org/x/crypto/bcrypt"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

func NewPostgresDb(db *sql.DB) *PostgresDBRepo {
	return &PostgresDBRepo{
		DB: db,
	}
}

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) Register(u models.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	db := m.DB
	_, err = db.Exec("INSERT INTO users (username, password, email, phone_num) VALUES ($1, $2, $3, $4)",
		u.Username, string(hashedPassword), u.Email, u.PhoneNum)
	if err != nil {
		return err
	}

	return nil
}

func (m *PostgresDBRepo) UserExistsWithEmail(email string) (bool, error) {
	dbInstance := m.DB

	var count int
	err := dbInstance.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (m *PostgresDBRepo) LoginUser(email, password string) (models.User, error) {
	var user models.User
	var hashedPassword string
	dbInstance := m.DB

	err := dbInstance.QueryRow("SELECT id, password, username, phone_num FROM users WHERE email = $1", email).Scan(&user.UserID, &hashedPassword, &user.Username, &user.PhoneNum)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found")
			return models.User{}, errors.New("user not found")
		}
		fmt.Printf("Error while querying database: %v\n", err)
		return models.User{}, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		fmt.Println("Invalid password")
		return models.User{}, errors.New("invalid password")
	}
	user.Password = password
	user.Email = email

	return user, nil
}

func (m *PostgresDBRepo) GetBuildings() ([]models.Building, error) {
	db := m.DB
	rows, err := db.Query("select b.id, b.description, b.address, b.country, b.guests_num, b.rooms_num, b.bathrooms_num, b.price_day, b.avalable_from, b.avalable_untill, b.user_id, b.imgurl, b.city,  c.name as category from buildings b inner join categories c on c.id = b.category_id")
	if err != nil {
		return nil, err
	}
	var buildingsArr []models.Building
	defer rows.Close()
	for rows.Next() {
		var id int
		var description string
		var address string
		var country string
		var guestsNum int
		var roomsNum int
		var bathroomsNum int
		var priceDay int
		var avalableFrom string
		var avalableUntill string
		var userId int
		var imgUrl string
		var city string
		var category string
		err = rows.Scan(&id, &description, &address, &country, &guestsNum, &roomsNum, &bathroomsNum, &priceDay, &avalableFrom, &avalableUntill, &userId, &imgUrl, &city, &category)

		if err != nil {
			// handle this error
			return nil, err
		}
		building := models.Building{
			Id:             id,
			Description:    description,
			Address:        address,
			Country:        country,
			GuestsNum:      guestsNum,
			RoomsNum:       roomsNum,
			BathroomsNum:   bathroomsNum,
			PriceDay:       priceDay,
			AvalableFrom:   avalableFrom,
			AvalableUntill: avalableUntill,
			UserId:         userId,
			ImgUrl:         imgUrl,
			City:           city,
			Category:       category,
		}
		buildingsArr = append(buildingsArr, building)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return buildingsArr, nil
}

func (m *PostgresDBRepo) InsertBuilding(building models.Building) (int, error) {
	db := m.DB
	query := `
        INSERT INTO buildings (
            description,
            address,
            country,
            category_id,
            guests_num,
            rooms_num,
            bathrooms_num,
            price_day,
            avalable_from,
            avalable_untill,
            user_id,
            imgurl,
            city
        )
        VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
        )
        RETURNING id;
    `

	var buildingID int
	err := db.QueryRow(
		query,
		building.Description,
		building.Address,
		building.Country,
		building.Category,
		building.GuestsNum,
		building.RoomsNum,
		building.BathroomsNum,
		building.PriceDay,
		building.AvalableFrom,
		building.AvalableUntill,
		building.UserId,
		building.ImgUrl,
		building.City,
	).Scan(&buildingID)

	if err != nil {
		return 0, err
	}

	return buildingID, nil
}

func (m *PostgresDBRepo) UpdateBuilding(building models.Building) error {
	db := m.DB
	query := `
		UPDATE buildings
		SET
			description = $2,
			address = $3,
			country = $4,
			category_id = $5,
			guests_num = $6,
			rooms_num = $7,
			bathrooms_num = $8,
			price_day = $9,
			avalable_from = $10,
			avalable_untill = $11,
			user_id = $12,
			imgurl = $13,
			city = $14
		WHERE id = $1;
	`

	_, err := db.Exec(query,
		building.Id,
		building.Description,
		building.Address,
		building.Country,
		building.Category,
		building.GuestsNum,
		building.RoomsNum,
		building.BathroomsNum,
		building.PriceDay,
		building.AvalableFrom,
		building.AvalableUntill,
		building.UserId,
		building.ImgUrl,
		building.City,
	)

	return err
}

func (m *PostgresDBRepo) DeleteBuilding(buildingID int) error {
	db := m.DB
	query := "DELETE FROM buildings WHERE id = $1"

	result, err := db.Exec(query, buildingID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no building found with the given ID")
	}

	return nil
}
