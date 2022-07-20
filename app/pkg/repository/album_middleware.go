// Package repository
// this file is not a middleware in the usual sense of this technology
package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// DeleteAllAlbums a helper function that is used to delete all albums from a certain artist,
// this function is also used when deleting an artist,
// so it is moved to the middleware
func DeleteAllAlbums(db *sqlx.DB, tx *sql.Tx, artistID int) error {
	queryGetAlbums := fmt.Sprintf("SELECT aa.album_id FROM %s aa WHERE aa.artist_id=$1", artistAlbumsTable)

	var indexes []int
	if err := db.Select(&indexes, queryGetAlbums, artistID); err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, id := range indexes {
		// First we need to delete all songs on this album
		if err := DeleteAllSongs(db, tx, id); err != nil {
			return err
		}

		queryDelAlbum := fmt.Sprintf("DELETE FROM %s al WHERE al.id=$1", albumsTable)
		if _, err := db.Exec(queryDelAlbum, id); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return nil
}

// CheckForAvailabilityInArtistAlbums a helper function,
// this function is used to check if an album is on an artist_albums table
func CheckForAvailabilityInArtistAlbums(db *sqlx.DB, tx *sql.Tx, artistID, albumID int) error {
	queryCheck := fmt.Sprintf("SELECT * FROM %s aa WHERE aa.artist_id=$1 AND aa.album_id=$2", artistAlbumsTable)
	a, errCheck := db.Exec(queryCheck, artistID, albumID)

	if errCheck != nil {
		_ = tx.Rollback()
		return errCheck
	}

	rAff, rAffErr := a.RowsAffected()
	if rAffErr != nil {
		return rAffErr
	}

	if rAff == 0 {
		return errors.New("this album or artist was deleted")
	}

	return nil
}
