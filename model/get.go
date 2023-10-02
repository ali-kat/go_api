package model

import "database/sql"

type (
	Tracks struct {
		Name         string         `json: "name"`
		Composer     sql.NullString `json: "composer"`
		Milliseconds int            `json: "duration"`
		Bytes        int            `json: "bytes"`
		UnitPrice    float64        `json: "price"`
		Album        string         `json: "album"`
		Genre        string         `json: "genre"`
		Format       string         `json: "format"`
	}
	Albums struct {
		Album    string `json: "title"`
		Composer string `json: "composer"`
	}
)
