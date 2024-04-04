package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func inc(num int) int {
	return num + 1
}

func handleIndex(count *Count) echo.HandlerFunc {
	return func(c echo.Context) error {
		count.Count = inc(count.Count)
		return c.Render(http.StatusOK, "index", count)
	}
}

func handleContact(page *Page) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "contact", page)
	}
}
