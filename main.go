package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func setupDB(db *sql.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS accounts (
		account_id TEXT PRIMARY KEY,
		balance REAL
	);`)

	db.Exec(`INSERT OR IGNORE INTO accounts (account_id, balance) VALUES ('A', 500.00);`)
	db.Exec(`INSERT OR IGNORE INTO accounts (account_id, balance) VALUES ('B', 200.00);`)

	fmt.Println("Accounts table ready.")
}

func main() {
	db, err := sql.Open("sqlite3", "bank.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	fmt.Println("Connected to database.")
	setupDB(db)

	// Begin a transaction
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		log.Fatal("Failed to begin transaction:", err)
	}

	// Execute transaction operations
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - 100 WHERE account_id = $1", "A")
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to credit Account B:", err)
	}

	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + 100 WHERE account_id = $1", "B")
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to credit Account B:", err)
	}

	// Check for sufficient balance
	var balance float64
	err = tx.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE account_id = $1", "A").Scan(&balance)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to check balance:", err)
	}
	if balance < 0 {
		tx.Rollback()
		log.Fatal("Insufficient balance in Account A")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to commit transaction:", err)
	}

	fmt.Println("Transaction compledted successfully.")

}
