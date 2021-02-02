package store

import (
	"database/sql"
	"strings"
	"tast_for_AvitoAdvertising/internal/model"
)

type AdvertRepository struct {
	store *Store
}

func (rep *AdvertRepository) CreateAdvert(advert *model.Advert) (*model.Advert, error) {
	if err := advert.Validate(); err != nil {
		return nil, err
	}
	if err := advert.ValidateLinks(); err != nil {
		return nil, err
	}

	if err := rep.store.db.QueryRow(
		"INSERT INTO adverts (name, description, photo_links, price) values ($1, $2, $3, $4) RETURNING id",
		advert.Name, advert.Description, advert.PhotoLinks, advert.Price,
	).Scan(&advert.AdvertId); err != nil {
		return nil, err
	}
	return advert, nil
}

func (rep *AdvertRepository) GetAdvertById(id int, optionalFields string) (*model.Advert, error) {
	advert := &model.Advert{}
	if err := rep.store.db.QueryRow(
		"SELECT * FROM adverts where id = $1",
		id).Scan(&advert.AdvertId, &advert.Name, &advert.Description, &advert.PhotoLinks,
		&advert.Price, &advert.CreationDate); err != nil {
		return nil, err
	}
	if !strings.Contains(optionalFields, "description") {
		advert.Description = ""
	}
	if !strings.Contains(optionalFields, "links") {
		advert.PhotoLinks = ""
	}
	return advert, nil
}

func (rep *AdvertRepository) GetAllAdverts(sort string) ([]model.Advert, error) {
	var rows *sql.Rows
	var err error
	if sort == "date" {
		rows, err = rep.store.db.Query("SELECT * FROM adverts order by creation_date")
	} else if sort == "-date" {
		rows, err = rep.store.db.Query("SELECT * FROM adverts order by creation_date desc")
	} else if sort == "price" {
		rows, err = rep.store.db.Query("SELECT * FROM adverts order by price")
	} else if sort == "-price" {
		rows, err = rep.store.db.Query("SELECT * FROM adverts order by price desc")
	} else {
		rows, err = rep.store.db.Query("SELECT * FROM adverts")
	}
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	adverts := make([]model.Advert, 0)

	for rows.Next() {
		newAdvert := model.Advert{}
		err = rows.Scan(&newAdvert.AdvertId, &newAdvert.Name, &newAdvert.Description, &newAdvert.PhotoLinks,
			&newAdvert.Price, &newAdvert.CreationDate)
		adverts = append(adverts, newAdvert)
	}
	return adverts, nil
}
