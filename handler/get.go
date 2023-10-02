package handler

import (
	"net/http"
	"server/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) FetchTracks(c echo.Context) (err error) {
	sqlStmt := `
		SELECT
			tracks.Name,
			Composer,
			Milliseconds,
			Bytes,
			UnitPrice,
			Title as Album,
			genres.Name as Genre,
			media_types.Name as Format
		FROM tracks 
		INNER JOIN albums USING (AlbumId) 
		INNER JOIN genres USING (GenreId)
		INNER JOIN media_types USING (MediaTypeId);
	`
	rows, err := h.DB.Query(sqlStmt)
	defer rows.Close()

	tracks := []model.Tracks{}
	for rows.Next() {
		t := model.Tracks{}
		err = rows.Scan(&t.Name, &t.Composer, &t.Milliseconds, &t.Bytes, &t.UnitPrice, &t.Album, &t.Genre, &t.Format)
		if err != nil {
			return err
		}
		tracks = append(tracks, t)
	}
	c.JSON(http.StatusOK, &tracks)
	return
}

func (h *Handler) FetchAlbums(c echo.Context) (err error) {
	sqlStmt := `
	SELECT 
		Title as Album, 
		Name as Composer 
	FROM albums 
	INNER JOIN artists USING (ArtistId);
	`

	rows, err := h.DB.Query(sqlStmt)
	defer rows.Close()

	albums := []model.Albums{}
	for rows.Next() {
		a := model.Albums{}

		err = rows.Scan(&a.Album, &a.Composer)
		if err != nil {
			return err
		}
		albums = append(albums, a)
	}
	c.JSON(http.StatusOK, &albums)
	return
}
