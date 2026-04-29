package main

import (
	"context"
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
		log.Fatal("Failed to created table:", err)
	}

	fmt.Println("Accounts table ready.")

}

func insertOperator(db *sql.DB, op_id string, rank string, flight string,
	ruck float64, jump int, agil_r float64, agil_l float64,
	deadlift int, pullups int, carry float64, shuttle float64,
	run float64) {

	_, err := db.Exec(`INSERT OR IGNORE INTO readiness (
						operator_id, rank, flight, ruck, long_jump,
						agility_r, agility_l, deadlift, pullups, carry,
						shuttle, run) VALUES ($1, $2, $3, $4, $5, $6, $7,
						$8, $9, $10, $11, $12);`, op_id, rank, flight, ruck,
		jump, agil_r, agil_l, deadlift, pullups, carry, shuttle,
		run)
	if err != nil {
		log.Fatal("Failed to insert operator")
	}

	fmt.Println("Operator succesfully inserted")
}

func updateOperator(db *sql.DB, op_id string, rank string, flight string,
	ruck float64, jump int, agil_r float64, agil_l float64,
	deadlift int, pullups int, carry float64, shuttle float64,
	run float64) {

	// Begin a transaction
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		log.Fatal("Failed to begin transaction:", err)
	}

	total := int(ruck) + jump + int(agil_r) + int(agil_l) + deadlift + pullups +
			int(carry) + int(shuttle) + int(run)
	status := "RED"
	if total > 500 {
		status = "GREEN"
	}

	// Execute transaction operations
	_, err = tx.ExecContext(ctx, `UPDATE readiness SET ruck = $1, long_jump = $2 ,
				agility_r = $3, agility_l = $4, deadlift = $5,
				pullups = $6, carry = $7, shuttle = $8, run = $9,
				total_score = $10, status = $11 WHERE operator_id = $12`,
			ruck, jump, agil_r, agil_l, deadlift, pullups, carry, shuttle,
			run, total, status, op_id)
	
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to update operator",err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to commit transaction:", err)
	}

	fmt.Println("Update Operator Transaction Completed Successfully")



}

func printOperator(db *sql.DB, op_id string) {
	var status string
	var total int

	err := db.QueryRow(`SELECT status, total_score FROM readiness WHERE operator_id = $1`, op_id).Scan(&status, &total)
	if err != nil {
		log.Fatal("Failed to query operator:", err)
	}

	fmt.Printf("Operator: %s | Score: %d | Status %s\n", op_id, total, status)
}

func main() {
	db, err := sql.Open("sqlite3", "readiness.db")
	if err != nil {
		log.Fatal("Faield to connect to database:", err)
	}
	defer db.Close()

	fmt.Println("Connected to database.")
	setupDB(db)
	insertOperator(db, "mike", "civ", "A", 20.0, 45, 4.0, 4.0, 200, 20, 20.0,
		65.0, 9.5)
	
	updateOperator(db, "mike", "civ", "A", 20.0, 45, 4.0, 4.0, 200, 20, 20.0,
		65.0, 9.5)
	
	printOperator(db, "mike")

	updateOperator(db, "mike", "civ", "A", 45.0, 95, 4.0, 4.0, 400, 20, 20.0,
		65.0, 9.5)
	printOperator(db, "mike")

}
