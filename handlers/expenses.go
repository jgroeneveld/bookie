package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jgroeneveld/bookie/db"
	"github.com/jgroeneveld/bookie/entities"
	"github.com/jgroeneveld/util"
)

type ExpensesHandler struct {
	DB *sql.DB
}

func (handler *ExpensesHandler) GetExpenses(resp http.ResponseWriter, req *http.Request) {
	expenses := db.GetExpenses(handler.DB)
	renderJSON(200, expenses, resp)
}

func (handler *ExpensesHandler) CreateExpense(resp http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	amountFloat, err := strconv.ParseFloat(params.Get("Amount"), 32)
	util.PanicIf(err)

	expense := entities.Expense{
		User:     entities.User(params.Get("User")),
		Category: entities.Category(params.Get("Category")),
		Amount:   entities.Money(amountFloat),
		SpentAt:  dateFromString(params.Get("Date")),
	}

	log.Println(params)
	log.Println(expense)

	db.InsertExpense(handler.DB, expense)

	renderJSON(201, "", resp)
}

func dateFromString(s string) time.Time {
	layout := "2006-01-02"

	d, _ := time.Parse(layout, s)

	return d
}

func renderJSON(status int, obj interface{}, resp http.ResponseWriter) {
	var result []byte
	var err error

	result, err = json.Marshal(obj)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	if _, err = io.Copy(resp, bytes.NewBuffer(result)); err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func ExpensesReport(render render.Render) {
// 	report := entities.ExpensesReport{
// 		MonthlyReports: []entities.MonthlyReport{
// 			{
// 				Month:       "2014-05",
// 				TotalAmount: 22.23,
// 				AmountByUsers: entities.UserMoneyMap{
// 					"Jaap":  12.23,
// 					"Hilke": 10.00,
// 				},
// 				AmountByCategory: entities.CategoryMoneyMap{
// 					"Edeka": 12.23,
// 					"Lidl":  10.00,
// 				},
// 			},
// 		},
// 	}
// 	render.JSON(200, report)
// }
