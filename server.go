package main

import "github.com/labstack/echo"

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.GET("/character/:id", func(c echo.Context) error {
		return c.String(200, "Character: "+c.Param("id"))
	})

	e.GET("/character", func(c echo.Context) error {
		return c.String(200, "All characters")
	})

	e.GET("/stage/:id", func(c echo.Context) error {
		return c.String(200, "Stage: "+c.Param("id"))
	})

	e.GET("/stage", func(c echo.Context) error {
		return c.String(200, "All stages")
	})

	return e.Start(s.listenAddr)
}
