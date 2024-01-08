package db

import (
	classes "renting/classes"
)

func GetBuildings() ([]classes.Building, error) {
	db := GetDBInstance()
	rows, err := db.Query("select b.id, b.description, b.address, b.country, b.guests_num, b.rooms_num, b.bathrooms_num, b.price_day, b.avalable_from, b.avalable_untill, b.user_id, b.imgurl, b.city,  c.name as category from buildings b inner join categories c on c.id = b.category_id")
	if err != nil {
		// handle this error better than this
		return nil, err
	}
	var buildingsArr []classes.Building
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
		building := classes.Building{
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
