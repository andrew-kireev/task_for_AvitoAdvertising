package model

type Advert struct {
	AdvertId     int    `json:"advert_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	PhotoLinks   string `json:"photo_links"`
	Price        int    `json:"price"`
	CreationDate string `json:"-"`
}
