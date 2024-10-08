package server

import (
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

func makeJsonRespFromName(name string, c echo.Context, db db.Storage) error {
	character, err := db.GetCharacterByName(name)
	if err != nil {
		return c.JSON(404, err)
	}
	return c.JSON(200, character)
}

func makeJsonRespFromId(id uint, c echo.Context, db db.Storage) error {
	character, err := db.GetCharacterByID(id)
	if err != nil {
		return c.JSON(404, err)
	}
	return c.JSON(200, character)
}
