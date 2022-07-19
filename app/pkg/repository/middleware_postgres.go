package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// DeleteAll a helper function that is used to delete all albums from a certain artist,
// this function is also used when deleting an artist,
// so it is moved to the middleware
func DeleteAll(db *sqlx.DB, artistID int) error {
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
