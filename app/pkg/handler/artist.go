package handler

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getAllArtistsResponse struct {
	Data []msh.Artist `json:"data"`
}

type getArtistByIDResponse struct {
	Artist msh.GetArtistWithAlbums `json:"artist"`
}

func (h *Handler) createArtist(c *gin.Context) {
	var input msh.Artist

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newArtist, err := h.services.Artist.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]any{
		"id":   newArtist.ID,
		"name": newArtist.Name,
		"age":  newArtist.Age,
	})
}

func (h *Handler) getArtists(c *gin.Context) {
	artists, err := h.services.GetAll()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllArtistsResponse{
		Data: artists,
	})
}

func (h *Handler) getArtistByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var artistInfo msh.GetArtistWithAlbums

	artistInfo, err = h.services.Artist.GetByID(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getArtistByIDResponse{
		Artist: artistInfo,
	})
}

func (h *Handler) updateArtist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input msh.UpdateArtistInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteArtist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err = h.services.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
