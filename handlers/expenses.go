package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jgroeneveld/bookie/db"
)

type ExpensesHandler struct {
	DB *sql.DB
}

func (handler *ExpensesHandler) GetExpenses(resp http.ResponseWriter, req *http.Request) {
	expenses := db.GetExpenses(handler.DB)
	renderJSON(200, expenses, resp)
}

func (handler *ExpensesHandler) CreateExpense(resp http.ResponseWriter, req *http.Request) {
	// TODO implement
	renderJSON(201, "", resp)
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
