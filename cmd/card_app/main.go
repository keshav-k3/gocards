package main

import (
	"card_app/internal/handler"
	"card_app/views/layouts"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pocketbase/pocketbase"
)

func main() {
	e := echo.New()

	pb := pocketbase.New()

	e.GET("/", func(c echo.Context) error {
		return handler.Render(c, http.StatusOK, layouts.BaseLayout())
	})

	go func() {
		if err := pb.Start(); err != nil {
			log.Printf("PocketBase error: %v", err)
		}
	}()

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
