package server

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Server) NewApi(e *echo.Echo) {
	char := e.Group("/character")
	{
		char.GET("", func(c echo.Context) error {
			characters, err := s.store.GetCharacters()
			if err != nil {
				return c.JSON(404, err)
			}
			return c.JSON(200, characters)
		})

		char.GET("/", func(c echo.Context) error {
			characters, err := s.store.GetCharacters()
			if err != nil {
				return c.JSON(404, err)
			}
			return c.JSON(200, characters)
		})

		char.GET("/:id", func(c echo.Context) error {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return makeJsonRespFromName(c.Param("id"), c, s.store)
			} else {
				return makeJsonRespFromId(uint(id), c, s.store)
			}
		})
	}

	song := e.Group("/song")
	{
		song.GET("", func(c echo.Context) error {
			songs, err := s.store.GetSongs()
			if err != nil {
				return c.JSON(404, err)
			}
			return c.JSON(200, songs)
		})

		song.GET("/", func(c echo.Context) error {
			songs, err := s.store.GetSongs()
			if err != nil {
				return c.JSON(404, err)
			}
			return c.JSON(200, songs)
		})

		song.GET("/:id", func(c echo.Context) error {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				song, err := s.store.GetSongByTitle("\"" + c.Param("id") + "\"")
				fmt.Println(c.Param("id"))
				if err != nil {
					return c.JSON(404, err)
				}
				return c.JSON(200, song)
			}

			song, err := s.store.GetSongByID(uint(id))
			if err != nil {
				return c.JSON(404, err)
			}
			return c.JSON(200, song)
		})
	}

	stage := e.Group("/stage")
	{
		stage.GET("", func(c echo.Context) error {
			return c.String(200, "All stages")
		})

		stage.GET("/", func(c echo.Context) error {
			return c.String(200, "All stages")
		})

		stage.GET("/:id", func(c echo.Context) error {
			return c.String(200, "Stage: "+c.Param("id"))
		})
	}
}
