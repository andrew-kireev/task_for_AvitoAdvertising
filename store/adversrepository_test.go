package store

import (
	"fmt"
	"reflect"
	"tast_for_AvitoAdvertising/internal/model"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestAdvertRepository_GetAdvertById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var elemId int = 1

	store := &Store{
		db:        db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "description", "photo_links", "price", "creation_date"})
	expect := []*model.Advert{
		{int(elemId), "some name", "some description",
			"link1 link2", 500, "2020-03-04"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.AdvertId, item.Name, item.Description, item.PhotoLinks, item.Price, item.CreationDate)
	}

	mock.ExpectQuery("SELECT").WithArgs(elemId).WillReturnRows(rows)
	item, err := store.Adverts().GetAdvertById(int(elemId), "description links")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	rows = sqlmock.NewRows([]string{"id", "name", "description", "photo_links", "price", "creation_date"}).
		AddRow(int(elemId), "some name", "",
			"", 500, "2020-03-04")

	mock.ExpectQuery("SELECT").WithArgs(elemId).WillReturnRows(rows)

	expected := &model.Advert{int(elemId), "some name",
		"", "",
		500, "2020-03-04",
	}

	item, err = store.Adverts().GetAdvertById(int(elemId), "")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expected) {
		t.Errorf("results not match, want %v, have %v", expected, item)
		return
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(elemId).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = store.Adverts().GetAdvertById(elemId, "")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

//func TestAdvertRepository_CreateAdvert(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("cant create mock: %s", err)
//	}
//	defer db.Close()
//
//	var elemId int = 1
//
//	store := &Store{
//		db:        db,
//	}
//
//	advert := &model.Advert{
//		elemId, "some name", "some description",
//		"link1 link2", 500, "2020-03-04",
//	}
//
//	mock.
//		ExpectExec(`INSERT INTO adverts`).
//		WithArgs(advert.Name, advert.Description, advert.PhotoLinks, advert.Price).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//
//
//	item, err := store.Adverts().CreateAdvert(advert)
//	if err != nil {
//		t.Errorf("unexpected err: %s", err)
//		return
//	}
//	if item.AdvertId != 1 {
//		t.Errorf("bad id: want %v, have %v", item.AdvertId, 1)
//		return
//	}
//
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}


func TestAdvertRepository_GetAllAdverts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var elemId int = 1

	store := &Store{
		db:        db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "description", "photo_links", "price", "creation_date"})
	expect := []model.Advert{
		{elemId, "some name", "some description",
			"link1 link2", 500, "2020-03-04"},
		{elemId, "some name", "some description",
			"link1 link2", 1000, "2019-03-04"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.AdvertId, item.Name, item.Description, item.PhotoLinks, item.Price, item.CreationDate)
	}

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	item, err := store.Adverts().GetAllAdverts("")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect) {
		t.Errorf("results not match, want %v, have %v", expect, item)
		return
	}

	rows = sqlmock.
		NewRows([]string{"id", "name", "description", "photo_links", "price", "creation_date"})

	for _, item := range expect {
		rows = rows.AddRow(item.AdvertId, item.Name, item.Description, item.PhotoLinks, item.Price, item.CreationDate)
	}
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	item, err = store.Adverts().GetAllAdverts("price")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect) {
		t.Errorf("results not match, want %v, have %v", expect, item)
		return
	}
}
