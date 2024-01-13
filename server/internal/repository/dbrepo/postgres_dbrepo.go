package dbrepo

import (
	"database/sql"
	"errors"
	"renting/models"

	"golang.org/x/crypto/bcrypt"
)

type PostgresDBRepo struct {
	DB *sql.DB
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

func (m *PostgresDBRepo) LoginUser(email, password string) (models.User, error) {
	var user models.User
	var hashedPassword string
	dbInstance := m.DB

	// query the database for the hashed password and admin flag based on the username
	err := dbInstance.QueryRow("SELECT user_id, password, email, phone_num FROM users WHERE username = $1", email).Scan(&user.UserID, &hashedPassword, &user.Email, &user.PhoneNum)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	// compare the hashed password from the database with the one the user provided.
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return models.User{}, errors.New("invalid password")
	}

	user.Password = password

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
