package model

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"strings"
)

type Advert struct {
	AdvertId     int    `json:"advert_id"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	PhotoLinks   string `json:"photo_links,omitempty"`
	Price        int    `json:"price"`
	CreationDate string `json:"-"`
}

func (adv *Advert) Validate() error {
	return validation.ValidateStruct(adv,
		validation.Field(&adv.Name, validation.Length(0, 10)),
		validation.Field(&adv.Description, validation.Length(0, 1000)))
}

func (adv *Advert) ValidateLinks() error {
	links := strings.Split(adv.PhotoLinks, " ")
	if len(links) > 3 {
		return errors.New("links field not valid")
	}
	return nil
}