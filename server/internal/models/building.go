package models

type Building struct {
	Id             int    `json:"id"`
	Description    string `json:"description"`
	Address        string `json:"address"`
	Country        string `json:"country"`
	Category       int `json:"category"`
	GuestsNum      int    `json:"guests_num"`
	RoomsNum       int    `json:"rooms_num"`
	BathroomsNum   int    `json:"bathrooms_num"`
	PriceDay       int    `json:"price_day"`
	AvalableFrom   string `json:"avalable_from"`
	AvalableUntill string `json:"avalable_untill"`
	UserId         int    `json:"user_id"`
	ImgUrl         string `json:"img_url"`
	City           string `json:"city"`
}
