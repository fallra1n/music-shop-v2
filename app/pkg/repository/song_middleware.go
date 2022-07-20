package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// CheckForAvailabilityInSongs a helper function,
// this function is used to check if a song is on an album_songs table
func CheckForAvailabilityInSongs(db *sqlx.DB, tx *sql.Tx, albumID, songID int) error {
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
