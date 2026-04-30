package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type EventScore struct {
	Name  string
	Score float64
}

func insertScoringRows(db *sql.DB, event string, rows []struct {
	min, max float64
	points   int
}) {
	for _, r := range rows {
		_, err := db.Exec(`INSERT OR IGNORE INTO scoring (event, min, max, points)
							VALUES ($1, $2, $3, $4)`, event, r.min, r.max, r.points)
		if err != nil {
			log.Fatal("Failed to insert", event, "scoring:", err)
		}
	}
	fmt.Println("Succesfully inserted", event, "rows")

}

func getPoints(db *sql.DB, event string, rawScore float64) int {
	var points int
	err := db.QueryRow(`SELECT points FROM scoring WHERE event = $1
						AND $2 >= min AND $2 <= max`, event, rawScore).Scan(&points)
	if err != nil {
		log.Fatal("Failed to get points for", event, err)
	}

	return points
}
