package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Operator struct {
	ID         string
	Rank       string
	Flight     string
	Status     string
	TotalScore int
}

func insertOperator(db *sql.DB, op Operator) {

	_, err := db.Exec(`INSERT OR IGNORE INTO readiness (
						operator_id, rank, flight) VALUES ($1, $2, $3);`, op.ID, op.Rank, op.Flight)
	if err != nil {
		log.Fatal("Failed to insert operator")
	}

	fmt.Println("Operator succesfully inserted")
}

func getOperatorsTotalScore(db *sql.DB) []Operator {
	operators := []Operator{}

	rows, err := db.Query("SELECT operator_id, status FROM readiness ORDER BY total_score DESC")
	if err != nil {
		log.Fatal("Failed to SELECT operators by total_score DESC", err)
	}
	defer rows.Close()

	for rows.Next() {
		var op Operator
		err := rows.Scan(&op.ID, &op.Status)
		if err != nil {
			log.Fatal("Failed rows.Scan", err)
		}
		operators = append(operators, op)
	}

	fmt.Println("GET Operators total_score DESC success")

	return operators
}

func getOperatorsFromFlight(db *sql.DB, flight string) []Operator {
	// make a list to return
	operators := []Operator{}

	// open a Query to the database
	rows, err := db.Query("SELECT operator_id, total_score FROM readiness WHERE flight = $1 ORDER BY total_score DESC", flight)
	if err != nil {
		log.Fatal("Query to getOperatorsFromFlight failed", err)
	}
	defer rows.Close()

	// now loop through query
	for rows.Next() {
		// create operator to append to list
		var op Operator
		err := rows.Scan(&op.ID, &op.TotalScore)
		if err != nil {
			log.Fatal("Failed rows.Scan", err)
		}
		operators = append(operators, op)
	}

	fmt.Println("Get Ops from Flight total_score DESC success")

	return operators
}

func updateOperatorFlight(db *sql.DB, opID string, flight string) {
	// updates an operators flight. Take operator ID and a new Flight as parameters
	_, err := db.Exec("UPDATE readiness SET flight = $1 WHERE operator_id = $2", flight, opID)
	if err != nil {
		log.Fatal("updateOperatorFlight failed", err)
	}

	fmt.Println("updateOperatorFlight success!")

}

func recordTest(db *sql.DB, op Operator, scores map[string]float64) {
	total := 0
	passed := true

	for key, value := range scores {
		points := getPoints(db, key, value)
		total += points

		if points == 0 {
			passed = false

		}
	}

	status := "GREEN"
	if !passed {
		status = "RED"
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		log.Fatal("Failed to begin recordTest transaction:", err)
	}

	_, err = tx.ExecContext(ctx, `UPDATE readiness SET ruck = $1, long_jump = $2,
							agility_r = $3, agility_l = $4, deadlift = $5, pullups = $6,
							carry = $7, shuttle = $8, run = $9, total_score = $10,
							status = $11 WHERE operator_id = $12`,
		scores["ruck"], scores["long_jump"], scores["agility_r"], scores["agility_l"],
		scores["deadlift"], scores["pullups"], scores["carry"], scores["shuttle"],
		scores["run"], total, status, op.ID)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to record test:", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to comit recordTest transaction:", err)
	}

	fmt.Println("Record Test succesful")

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
		log.Fatal("Failed to update operator", err)
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
