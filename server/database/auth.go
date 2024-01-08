package db

import (
	"database/sql"
	"errors"
	classes "renting/classes"

	"golang.org/x/crypto/bcrypt"
)

func Register(username, email, password, phoneNum string) error {

	// Hash the password before storing it in your database.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Insert the new user into the database.
	db := GetDBInstance()
	_, err = db.Exec("INSERT INTO users (username, password, email, phone_num) VALUES ($1, $2, $3, $4)",
		username, string(hashedPassword), email, phoneNum)
	if err != nil {
		return err
	}

	return nil
}

func LoginUser(username, password string) (classes.User, error) {
	var user classes.User
	var hashedPassword string
	dbInstance := GetDBInstance()

	// query the database for the hashed password and admin flag based on the username
	err := dbInstance.QueryRow("SELECT user_id, password, email, phone_num FROM users WHERE username = $1", username).Scan(&user.UserID, &hashedPassword, &user.Email, &user.PhoneNum)
	if err != nil {
		if err == sql.ErrNoRows {
			return classes.User{}, errors.New("user not found")
		}
		return classes.User{}, err
	}

	// compare the hashed password from the database with the one the user provided.
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return classes.User{}, errors.New("invalid password")
	}

	user.Username = username
	user.Password = password

	return user, nil
}
