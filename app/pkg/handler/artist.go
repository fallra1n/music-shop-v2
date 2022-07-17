package handler

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createArtist(c *gin.Context) {
	var input msh.Artist

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newArtist, err := h.services.CreateArtist(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id":   newArtist.ID,
		"name": newArtist.Name,
		"age":  newArtist.Age,
	})

}

func (h *Handler) getArtists(c *gin.Context) {
}

func (h *Handler) getArtistByID(c *gin.Context) {
}

func (h *Handler) updateArtist(c *gin.Context) {
}

func (h *Handler) deleteArtist(c *gin.Context) {
}
