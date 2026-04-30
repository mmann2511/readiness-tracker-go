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

	fmt.Println("Accounts table ready.")

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

	fmt.Println("Scoring table ready.")
}
