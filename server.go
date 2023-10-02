package main

import (
	"database/sql"
	"fmt"
	"server/handler"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

	db, err := sql.Open("sqlite3", "./chinook.db")
	if err != nil {
		fmt.Println("ERROR")
	}
	defer db.Close()

	h := &handler.Handler{DB: db}

	e.GET("/albums", h.FetchAlbums)
	e.GET("/tracks", h.FetchTracks)

	e.Logger.Fatal(e.Start(":8090"))
}
