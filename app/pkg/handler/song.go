package handler

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllSongsResponse struct {
	Data []msh.Song `json:"data"`
}

func (h *Handler) createSong(c *gin.Context) {
	if _, err := CheckID(c, "id"); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	albumID, err := CheckID(c, "album_id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input msh.Song
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var newSong msh.Song
	newSong, err = h.services.Song.Create(albumID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]any{
		"id":    newSong.ID,
		"title": newSong.Title,
		"text":  newSong.Text,
		"album": newSong.Album,
	})
}

func (h *Handler) getSongs(c *gin.Context) {
	if _, err := CheckID(c, "id"); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	albumID, err := CheckID(c, "album_id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var songs []msh.Song
	songs, err = h.services.Song.GetAll(albumID)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllSongsResponse{
		Data: songs,
	})
}

func (h *Handler) getSongByID(c *gin.Context) {
	_, albumID, songID, err := CheckAllID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var song msh.GetSongOutput
	if song, err = h.services.Song.GetByID(albumID, songID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"title": song.Title,
		"text":  song.Text,
		"album": song.Album,
	})
}

func (h *Handler) updateSong(c *gin.Context) {
	_, albumID, songID, err := CheckAllID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input msh.UpdateSongInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Song.Update(albumID, songID, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteSong(c *gin.Context) {
	_, albumID, songID, err := CheckAllID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Song.Delete(albumID, songID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteAllSongs(c *gin.Context) {
	albumID, err := CheckID(c, "album_id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Song.DeleteAll(albumID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
