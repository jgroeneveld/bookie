package db

import (
	"database/sql"
	"io/ioutil"
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

func Migrate(db *sql.DB) {
	buffer, err := ioutil.ReadFile("schema.sql")
	util.PanicIf(err)
	schema := string(buffer)

	_, err = db.Exec(schema)
	util.PanicIf(err)
}

func GetExpense(db *sql.DB, id int) (error, *entities.Expense) {
	query := "select " +
		"category, amount, created_at, spent_at, username " +
		"from expenses " +
		"where id = $1 " +
		"limit 1"
 
	row := db.QueryRow(query, id)

	var category string
	var amount float32
	var createdAt time.Time
	var spentAt time.Time
	var user string

	if err := row.Scan(&category, &amount, &createdAt, &spentAt, &user); err != nil {
		return err, nil
	}

	expense := &entities.Expense{
		ID:  id,
		Category:  entities.Category(category),
		Amount:    entities.Money(amount),
		CreatedAt: createdAt,
		SpentAt:   spentAt,
		User:      entities.User(user),
	}

	return nil, expense
}

func GetExpenses(db *sql.DB) (error, []entities.Expense) {
	expenses := []entities.Expense{}

	rows, err := db.Query("select " +
		"id, category, amount, created_at, spent_at, username " +
		"from expenses " +
		"order by spent_at desc, id desc limit 1000")
	defer rows.Close()

	if err != nil {
		return err, nil
	}

	for rows.Next() {
		var id int
		var category string
		var amount float32
		var createdAt time.Time
		var spentAt time.Time
		var user string

		if err = rows.Scan(&id, &category, &amount, &createdAt, &spentAt, &user); err != nil {
			return err, nil
		}

		expense := entities.Expense{
			ID:  id,
			Category:  entities.Category(category),
			Amount:    entities.Money(amount),
			CreatedAt: createdAt,
			SpentAt:   spentAt,
			User:      entities.User(user),
		}

		expenses = append(expenses, expense)
	}

	return nil, expenses
}

func InsertExpense(db *sql.DB, expense entities.Expense) error {
	spentAt := expense.SpentAt.Format("2006-01-02")
	log.Println("spentAt", spentAt)

	result, err := db.Exec("insert into expenses "+
		"(username, category, amount, spent_at)"+
		" VALUES "+
		"($1, $2, $3, $4)",
		string(expense.User), string(expense.Category), int(expense.Amount), spentAt)

	log.Println("result", result)
	log.Println("err", err)
	return err
}

func GetExpensesReport(db *sql.DB) (error, entities.ExpensesReport) {
	report := entities.ExpensesReport{
		AmountByUsers: entities.UserMoneyMap{
			"Jaap":  GetTotalAmountForUser(db, "Jaap"),
			"Hilke": GetTotalAmountForUser(db, "Hilke"),
		},
		MonthlyReports: []entities.MonthlyReport{},
	}

	return nil, report
}

func GetTotalAmountForUser(db *sql.DB, user entities.User) entities.Money {
	username := string(user)
	row := db.QueryRow("select sum(amount) from expenses where username = $1", username)
	var amount float32
	row.Scan(&amount)
	return entities.Money(amount)
}
