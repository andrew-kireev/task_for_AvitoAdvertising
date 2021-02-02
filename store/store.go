package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
	advertRep *AdvertRepository
	config *Config
}

func NewStore(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (store *Store) Open() error {
	db, err := sql.Open("postgres", store.config.DataBaseUrl)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	store.db = db

	return nil
}

func (store *Store) Close() {

}

func (store *Store) Adverts() *AdvertRepository {
	if store.advertRep != nil {
		return store.advertRep
	}

	store.advertRep = &AdvertRepository{
		store: store,
	}
	return store.advertRep
}