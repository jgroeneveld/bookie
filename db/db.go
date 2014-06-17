package db

import (
	"database/sql"
	"os"
	"time"

	"github.com/jgroeneveld/bookie/entities"
	"github.com/jgroeneveld/util"
	_ "github.com/lib/pq"
)

func OpenDb() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	util.PanicIf(err)

	return db
}

func GetExpenses(db *sql.DB) []entities.Expense {
	expenses := []entities.Expense{}

	rows, err := db.Query("select category, amount, created_at, username from expenses")
	util.PanicIf(err)
	defer rows.Close()

	for rows.Next() {
		var category string
		var amount float32
		var createdAt time.Time
		var user string

		if err = rows.Scan(&category, &amount, &createdAt, &user); err != nil {
			util.PanicIf(err)
		}

		expense := entities.Expense{
			Category:  entities.Category(category),
			Amount:    entities.Money(amount),
			CreatedAt: createdAt,
			User:      entities.User(user),
		}

		expenses = append(expenses, expense)
	}

	return expenses
}
