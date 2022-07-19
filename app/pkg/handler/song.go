package handler

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createSong(c *gin.Context) {
	if _, err := strconv.Atoi(c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	albumID, err := strconv.Atoi(c.Param("album_id"))
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

	c.JSON(http.StatusOK, map[string]any{
		"id":    newSong.ID,
		"title": newSong.Title,
		"text":  newSong.Text,
		"album": newSong.Album,
	})
}

func (h *Handler) getSongs(c *gin.Context) {
}

func (h *Handler) getSongByID(c *gin.Context) {
}

func (h *Handler) updateSong(c *gin.Context) {
}

func (h *Handler) deleteSong(c *gin.Context) {
}
