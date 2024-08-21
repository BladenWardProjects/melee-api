package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	// "github.com/stretchr/testify/assert"
)

var greeting = "Welcome to the Super Smash Bros. Melee API!"

func TestGreeting(t *testing.T) {
	s := &Server{}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(greeting))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	s.Greeting(c)
	got := rec.Body.String()
	want := greeting

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
