// Package repository
// this file is not a middleware in the usual sense of this technology
package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// CheckForAvailabilityInSongs a helper function,
// this function is used to check if a song is on an album_songs table
func CheckForAvailabilityInAlbumSongs(db *sqlx.DB, tx *sql.Tx, albumID, songID int) error {
	queryCheck := fmt.Sprintf("SELECT * FROM %s ast WHERE ast.album_id=$1 AND ast.song_id=$2", albumSongsTable)
	a, errCheck := db.Exec(queryCheck, albumID, songID)
	if errCheck != nil {
		_ = tx.Rollback()
		return errCheck
	}

	rAff, errAff := a.RowsAffected()
	if errAff != nil {
		return errAff
	}

	if rAff == 0 {
		return errors.New("this song or album was deleted")
	}

	return nil
}

func CheckForAvailabilityInAlbums(db *sqlx.DB, tx *sql.Tx, albumID int) error {
	queryCheck := fmt.Sprintf("SELECT * FROM %s al WHERE al.id=$1", albumsTable)
	a, errCheck := db.Exec(queryCheck, albumID)
	if errCheck != nil {
		_ = tx.Rollback()
		return errCheck
	}

	rAff, errAff := a.RowsAffected()
	if errAff != nil {
		return errAff
	}

	if rAff == 0 {
		return errors.New("this album does not exist")
	}

	return nil
}

// DeleteAllSongs a helper function that is used to delete all songs from a certain album,
// this function is also used when deleting an album,
// so it is moved to the middleware
func DeleteAllSongs(db *sqlx.DB, tx *sql.Tx, albumID int) error {
	queryGetSongs := fmt.Sprintf("SELECT st.song_id FROM %s st WHERE st.album_id=$1", albumSongsTable)

	var indexes []int
	if err := db.Select(&indexes, queryGetSongs, albumID); err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, id := range indexes {
		queryDelSong := fmt.Sprintf("DELETE FROM %s st WHERE st.id=$1", songsTable)
		if _, err := db.Exec(queryDelSong, id); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return nil
}
