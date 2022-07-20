package handler

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createAlbum(c *gin.Context) {
	artistID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input msh.Album
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var newAlbum msh.Album
	newAlbum, err = h.services.Album.Create(artistID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]any{
		"id":     newAlbum.ID,
		"title":  newAlbum.Title,
		"price":  newAlbum.Price,
		"artist": newAlbum.Artist,
		"date":   newAlbum.Date,
	})
}

func (h *Handler) getAlbumByID(c *gin.Context) {
	artistID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var albumID int
	albumID, err = strconv.Atoi(c.Param("album_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var album msh.GetAlbumOutput
	album, err = h.services.Album.GetByID(artistID, albumID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"title":  album.Title,
		"price":  album.Price,
		"artist": album.Artist,
		"date":   album.Date,
	})
}

func (h *Handler) updateAlbum(c *gin.Context) {
	albumID, err := strconv.Atoi(c.Param("album_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input msh.UpdateAlbumInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Album.Update(albumID, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteAllAlbums(c *gin.Context) {
	artistID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DeleteAll(artistID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteAlbum(c *gin.Context) {
	albumID, err := strconv.Atoi(c.Param("album_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Album.Delete(albumID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
