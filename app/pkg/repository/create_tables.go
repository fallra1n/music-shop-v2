package repository

import (
	"github.com/jmoiron/sqlx"
)

type CreateTables struct {
	db *sqlx.DB
}

func NewCrateTables(db *sqlx.DB) *CreateTables {
	return &CreateTables{db: db}
}

func (ct *CreateTables) CreateAllTables() error {
	if err := ct.CreateAlbumsSchema(); err != nil {
		return err
	}
	if err := ct.CreateArtistsSchema(); err != nil {
		return err
	}
	if err := ct.CreateSongsSchema(); err != nil {
		return err
	}
	if err := ct.CreateArtistAlbumSchema(); err != nil {
		return err
	}
	if err := ct.CreateAlbumSongSchema(); err != nil {
		return err
	}

	return nil
}

func (ct *CreateTables) CreateArtistsSchema() error {
	query := "CREATE TABLE IF NOT EXISTS artists(" +
		"id serial primary key," +
		"name varchar(255) not null," +
		"age integer not null" +
		")"

	if _, err := ct.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (ct *CreateTables) CreateAlbumsSchema() error {
	query := "CREATE TABLE IF NOT EXISTS albums(" +
		"id serial primary key," +
		"title varchar(255) not null," +
		"price decimal not null," +
		"artist varchar(255) not null," +
		"date varchar(255) not null" +
		")"

	if _, err := ct.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (ct *CreateTables) CreateSongsSchema() error {
	query := "CREATE TABLE IF NOT EXISTS songs(" +
		"id serial primary key," +
		"title varchar(255) not null," +
		"text Text not null," +
		"album varchar(255) not null" +
		")"

	if _, err := ct.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (ct *CreateTables) CreateArtistAlbumSchema() error {
	query := "CREATE TABLE IF NOT EXISTS artist_albums(" +
		"id serial primary key," +
		"artist_id integer references artists(id) on delete cascade not null," +
		"album_id integer references albums(id) on delete cascade not null" +
		")"

	if _, err := ct.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (ct *CreateTables) CreateAlbumSongSchema() error {
	query := "CREATE TABLE IF NOT EXISTS album_songs(" +
		"id serial primary key," +
		"album_id integer references albums(id) on delete cascade not null," +
		"song_id integer references songs(id) on delete cascade not null" +
		")"

	if _, err := ct.db.Exec(query); err != nil {
		return err
	}
	return nil
}
