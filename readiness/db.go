package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func setupDB(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS readiness (
		operator_id TEXT PRIMARY KEY,
		rank TEXT,
		flight TEXT,
		status TEXT, 
		total_score INTEGER,
		ruck REAL,
		long_jump INTEGER,
		agility_r REAL,
		agility_l REAL,
		deadlift INTEGER,
		pullups INTEGER,
		carry REAL,
		shuttle REAL,
		run REAL

	);`)
	if err != nil {
		log.Fatal("Failed to create accounts table:", err)
	}
	createScoringTable(db)
	seedDB(db)

	fmt.Println("Readiness table ready.")

}

func createScoringTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS scoring (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event TEXT,
		min REAL,
		max REAL,
		points INTEGER
	)`)
	if err != nil {
		log.Fatal("Failed to create scoring table:", err)
	}

	insertScoringRows(db, "ruck", []struct {
		min, max float64
		points   int
	}{
		{0, 49, 20},
		{50, 999, 0},
	})

	insertScoringRows(db, "agility_r", []struct {
		min, max float64
		points   int
	}{
		{0.0, 4.99, 5},
		{5.0, 5.24, 4},
		{5.23, 5.50, 3},
		{5.51, 9.99, 0},
	})

	insertScoringRows(db, "agility_l", []struct {
		min, max float64
		points   int
	}{
		{0.0, 4.99, 5},
		{5.0, 5.24, 4},
		{5.23, 5.50, 3},
		{5.51, 9.99, 0},
	})

	insertScoringRows(db, "long_jump", []struct {
		min, max float64
		points   int
	}{
		{94, 999, 10},
		{85, 93, 9},
		{76, 84, 8},
		{0, 75, 0},
	})

	insertScoringRows(db, "deadlift", []struct {
		min, max float64
		points   int
	}{
		{360, 999, 10},
		{325, 359, 9},
		{305, 324, 8},
		{270, 304, 7},
		{0, 269, 0},
	})

	insertScoringRows(db, "pullups", []struct {
		min, max float64
		points   int
	}{
		{16, 30, 10},
		{15, 15, 9},
		{14, 14, 8},
		{12, 13, 7},
		{10, 11, 6},
		{0, 9, 0},
	})

	insertScoringRows(db, "carry", []struct {
		min, max float64
		points   int
	}{
		{0, 21, 10},
		{22, 24, 9},
		{25, 27, 8},
		{28, 29, 7},
		{30, 999, 0},
	})

	insertScoringRows(db, "shuttle", []struct {
		min, max float64
		points   int
	}{
		{0, 67.7, 10},
		{67.8, 71.1, 9},
		{71.2, 80.5, 8},
		{80.6, 999.9, 0},
	})

	insertScoringRows(db, "run", []struct {
		min, max float64
		points   int
	}{
		{0, 9, 20},
		{9.5, 10.25, 19},
		{10.26, 11.25, 18},
		{11.26, 12.5, 17},
		{12.6, 20.99, 0},
	})

	fmt.Println("Scoring table ready.")
}

func seedDB(db *sql.DB) {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM readiness").Scan(&count)
	if count == 0 {
		// insert test data
		insertOperator(db, Operator{ID: "mike", Rank: "civ", Flight: "A"})
		insertOperator(db, Operator{ID: "fred", Rank: "1LT", Flight: "B"})
		insertOperator(db, Operator{ID: "george", Rank: "2LT", Flight: "C"})
		insertOperator(db, Operator{ID: "ginny", Rank: "CAP", Flight: "A"})
		insertOperator(db, Operator{ID: "ron", Rank: "civ", Flight: "B"})
		insertOperator(db, Operator{ID: "draco", Rank: "MAJ", Flight: "C"})
		insertOperator(db, Operator{ID: "herminoe", Rank: "civ", Flight: "A"})
		insertOperator(db, Operator{ID: "luna", Rank: "COL", Flight: "B"})
		insertOperator(db, Operator{ID: "tonks", Rank: "1LT", Flight: "C"})
		insertOperator(db, Operator{ID: "lupin", Rank: "2LT", Flight: "A"})
	}
}
