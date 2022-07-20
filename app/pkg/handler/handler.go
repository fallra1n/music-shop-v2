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

	api := router.Group("/api")
	{
		artists := api.Group("/artists")
		{
			artists.POST("/", h.createArtist)      // done
			artists.GET("/", h.getArtists)         // done
			artists.GET("/:id", h.getArtistByID)   // done
			artists.PUT("/:id", h.updateArtist)    // done
			artists.DELETE("/:id", h.deleteArtist) // done

			albums := artists.Group("/:id/albums")
			{
				albums.POST("/", h.createAlbum)            // done
				albums.GET("/:album_id", h.getAlbumByID)   // done
				albums.PUT("/:album_id", h.updateAlbum)    // done
				albums.DELETE("/", h.deleteAllAlbums)      // done
				albums.DELETE("/:album_id", h.deleteAlbum) // done

				songs := albums.Group("/:album_id/songs")
				{
					songs.POST("/", h.createSong)         // done
					songs.GET("/", h.getSongs)            // done
					songs.GET("/:song_id", h.getSongByID) // done
					songs.PUT("/:song_id", h.updateSong)
					songs.DELETE("/", h.deleteAllSongs)     // done
					songs.DELETE("/:song_id", h.deleteSong) // done
				}
			}
		}
	}
	return router
}
