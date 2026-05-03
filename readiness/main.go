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

	// op := Operator{ID: "george"}

	// recordTest(db, op, map[string]float64{
	// 	"ruck":      40.0,
	// 	"long_jump": 94,
	// 	"agility_r": 4.0,
	// 	"agility_l": 4.0,
	// 	"deadlift":  400,
	// 	"pullups":   20,
	// 	"carry":     20,
	// 	"shuttle":   60.0,
	// 	"run":       9.0,
	// })

	// op2 := Operator{ID: "draco"}

	// recordTest(db, op2, map[string]float64{
	// 	"ruck":      40.0,
	// 	"long_jump": 85,
	// 	"agility_r": 5.0,
	// 	"agility_l": 5.0,
	// 	"deadlift":  350,
	// 	"pullups":   14,
	// 	"carry":     25,
	// 	"shuttle":   70.0,
	// 	"run":       10.0,
	// })

	// op3 := Operator{ID: "tonks"}

	// recordTest(db, op3, map[string]float64{
	// 	"ruck":      50.0,
	// 	"long_jump": 94,
	// 	"agility_r": 4.0,
	// 	"agility_l": 4.0,
	// 	"deadlift":  400,
	// 	"pullups":   20,
	// 	"carry":     20,
	// 	"shuttle":   60.0,
	// 	"run":       9.0,
	// })

	// operators := getOperatorsFromFlight(db, "C")

	// for _, op := range operators {
	// 	fmt.Printf("Operator: %s, Total Score: %d\n", op.ID, op.TotalScore)
	// }

	op4 := Operator{ID: "mike"}
	recordTest(db, op4, map[string]float64{
		"ruck":      50.0,
		"long_jump": 94,
		"agility_r": 4.0,
		"agility_l": 4.0,
		"deadlift":  400,
		"pullups":   20,
		"carry":     20,
		"shuttle":   60.0,
		"run":       9.0,
	})

	op5 := Operator{ID: "fred"}
	recordTest(db, op5, map[string]float64{
		"ruck":      50.0,
		"long_jump": 94,
		"agility_r": 4.0,
		"agility_l": 4.0,
		"deadlift":  400,
		"pullups":   20,
		"carry":     20,
		"shuttle":   60.0,
		"run":       9.0,
	})

	operators := getOperatorsFromFlight(db, "C")
	for _, op := range operators {
		fmt.Printf("Operator %s | Score: %d\n", op.ID, op.TotalScore)
	}

	updateOperatorFlight(db, "fred", "C")

	operators2 := getOperatorsFromFlight(db, "C")
	for _, op := range operators2 {
		fmt.Printf("Operator %s | Score: %d\n", op.ID, op.TotalScore)
	}

}
