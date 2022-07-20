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
func DeleteAllAlbums(db *sqlx.DB, artistID int) error {
	queryGetAlbums := fmt.Sprintf("SELECT aa.album_id FROM %s aa WHERE aa.artist_id=$1", artistAlbumsTable)

	var indexes []int
	if err := db.Select(&indexes, queryGetAlbums, artistID); err != nil {
		return err
	}

	for _, id := range indexes {
		queryDelAlbum := fmt.Sprintf("DELETE FROM %s al WHERE al.id=$1", albumsTable)
		if _, err := db.Exec(queryDelAlbum, id); err != nil {
			return err
		}
	}

	return nil
}

// CheckForAvailabilityInAlbums a helper function,
// this function is used to check if a album is on an artist_albums table
func CheckForAvailabilityInAlbums(db *sqlx.DB, tx *sql.Tx, artistID, albumID int) error {
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
