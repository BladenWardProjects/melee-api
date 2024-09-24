package server

import (
	"strconv"

	"github.com/BladenWard/melee-api/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	ListenAddr string
	store      db.Storage
}

func NewServer(ListenAddr string, store db.Storage) *Server {
	return &Server{
		ListenAddr: ListenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	e := echo.New()

	// e.Pre(middleware.AddTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())

	e.GET("/", s.Greeting)

	s.NewApi(e)

	return e.Start(s.ListenAddr)
}

func (s *Server) Greeting(c echo.Context) error {
	return c.String(200, "Welcome to the Super Smash Bros. Melee API!")
}

func (s *Server) NewApi(e *echo.Echo) {
	char := e.Group("/character")
	{
		char.GET("", func(c echo.Context) error {
			return c.String(200, "All characters")
		})

		char.GET("/", func(c echo.Context) error {
			return c.String(200, "All characters")
		})

		char.GET("/:id", func(c echo.Context) error {
			id, _ := strconv.Atoi(c.Param("id"))
			character, err := s.store.GetCharacterByID(uint(id))
			if err != nil {
				return c.JSON(404, err)
			}
			return c.JSON(200, character)
		})

		char.GET("/name/:name", func(c echo.Context) error {
			retrievedChar, err := s.store.GetCharacterByName(c.Param("name"))
			if err != nil {
				return c.JSON(404, err)
			}
			return c.JSON(200, retrievedChar)
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
