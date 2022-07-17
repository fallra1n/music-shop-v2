package handler

import (
	"github.com/asssswv/music-shop-v2/app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/", getHomePage)

	api := router.Group("/api")
	{
		artists := api.Group("/artists")
		{
			artists.POST("/", h.createArtist)
			artists.GET("/", h.getArtists)
			artists.GET("/:id", h.getArtistByID)
			artists.PUT("/:id", h.updateArtist)
			artists.DELETE("/:id", h.deleteArtist)
		}

		albums := api.Group("/albums")
		{
			albums.POST("/", h.createAlbum)
			albums.GET("/", h.getAlbums)
			albums.GET("/:id", h.getAlbumByID)
			albums.PUT("/:id", h.updateAlbum)
			albums.DELETE("/:id", h.deleteAlbum)

			songs := albums.Group("/:id/items")
			{
				songs.POST("/", h.createSong)
				songs.GET("/", h.getSongs)
				songs.GET("/:song_id", h.getSongByID)
				songs.PUT("/:song_id", h.updateSong)
				songs.DELETE("/:song_id", h.deleteSong)
			}
		}
	}
	return router
}
