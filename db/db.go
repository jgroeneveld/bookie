package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jgroeneveld/bookie/entities"
	"github.com/jgroeneveld/util"
	_ "github.com/lib/pq"
)

func OpenDb() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	if len(url) == 0 {
		log.Fatal("DATABASE_URL missing")
	}

	db, err := sql.Open("postgres", url)
	util.PanicIf(err)

	return db
}

func GetExpenses(db *sql.DB) []entities.Expense {
	expenses := []entities.Expense{}

	rows, err := db.Query("select " +
		"category, amount, created_at, spent_at, username " +
		"from expenses " +
		"order by spent_at desc limit 1000")
	util.PanicIf(err)
	defer rows.Close()

	for rows.Next() {
		var category string
		var amount float32
		var createdAt time.Time
		var spentAt time.Time
		var user string

		if err = rows.Scan(&category, &amount, &createdAt, &spentAt, &user); err != nil {
			util.PanicIf(err)
		}

		expense := entities.Expense{
			Category:  entities.Category(category),
			Amount:    entities.Money(amount),
			CreatedAt: createdAt,
			SpentAt:   spentAt,
			User:      entities.User(user),
		}

		expenses = append(expenses, expense)
	}

	return expenses
}

func InsertExpense(db *sql.DB, expense entities.Expense) error {
	spentAt := expense.SpentAt.Format("2006-01-02")
	log.Println("spentAt", spentAt)

	result, err := db.Exec(fmt.Sprintf("insert into expenses "+
		"(username, category, amount, spent_at)"+
		" VALUES "+
		"('%s', '%s', %f, '%s')",
		expense.User, expense.Category, expense.Amount, spentAt))

	log.Println("result", result)
	log.Println("err", err)
	return nil
}
