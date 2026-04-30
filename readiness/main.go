package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "readiness.db")
	if err != nil {
		log.Fatal("Faield to connect to database:", err)
	}
	defer db.Close()

	fmt.Println("Connected to database.")
	setupDB(db)
	createScoringTable(db)
	insertScoringRows(db, "long_jump", []struct {
		min, max float64
		points   int
	}{
		{94, 999, 10},
		{85, 93, 9},
		{76, 84, 8},
		{0, 75, 0},
	})

	op := Operator{ID: "mike", Rank: "civ", Flight: "A"}
	insertOperator(db, op)
	recordTest(db, op, map[string]float64{
		"long_jump": 94,
	})
	printOperator(db, "mike")

}
