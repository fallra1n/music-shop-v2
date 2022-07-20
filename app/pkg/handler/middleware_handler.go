// Package handler
// this file is not a middleware in the usual sense of this technology
package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CheckID(c *gin.Context, key string) (int, error) {
	return strconv.Atoi(c.Param(key))
}

func CheckAllID(c *gin.Context) (int, int, int, error) {
	artistID, err := CheckID(c, "id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return 0, 0, 0, err
	}

	var albumID int
	if albumID, err = CheckID(c, "album_id"); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return 0, 0, 0, err
	}

	var songID int
	if songID, err = CheckID(c, "song_id"); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return 0, 0, 0, err
	}

	return artistID, albumID, songID, nil
}
