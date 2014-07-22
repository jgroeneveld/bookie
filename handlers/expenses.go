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

	"github.com/gorilla/mux"
	"github.com/jgroeneveld/bookie/db"
	"github.com/jgroeneveld/bookie/entities"
	"github.com/jgroeneveld/util"
)

type ExpensesHandler struct {
	DB *sql.DB
}

func (handler *ExpensesHandler) GetExpense(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	util.PanicIf(err)

	err, expenses := db.GetExpense(handler.DB, id)
	if err != nil {
		renderInternalError(resp, err)
	} else {
		renderJSON(resp, 200, expenses)
	}
}

func (handler *ExpensesHandler) GetExpenses(resp http.ResponseWriter, req *http.Request) {
	err, expenses := db.GetExpenses(handler.DB)
	if err != nil {
		renderInternalError(resp, err)
	} else {
		renderJSON(resp, 200, expenses)
	}
}

func (handler *ExpensesHandler) CreateExpense(resp http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	amount, err := strconv.Atoi(params.Get("Amount"))
	if err != nil {
		http.Error(resp, "Amount invalid", http.StatusBadRequest)
		return
	}

	expense := entities.Expense{
		User:     entities.User(params.Get("User")),
		Category: entities.Category(params.Get("Category")),
		Amount:   entities.Money(amount),
		SpentAt:  dateFromString(params.Get("Date")),
	}

	log.Println(params)
	log.Println(expense)

	err = db.InsertExpense(handler.DB, expense)

	if err != nil {
		renderInternalError(resp, err)
	} else {
		renderJSON(resp, 201, expense)
	}
}

func (handler *ExpensesHandler) GetExpensesReport(resp http.ResponseWriter, req *http.Request) {
	err, report := db.GetExpensesReport(handler.DB)

	if err != nil {
		renderInternalError(resp, err)
	} else {
		renderJSON(resp, 200, report)
	}
}

func dateFromString(s string) time.Time {
	layout := "2006-01-02"

	d, _ := time.Parse(layout, s)

	return d
}

func renderInternalError(resp http.ResponseWriter, err error) {
	http.Error(resp, err.Error(), http.StatusInternalServerError)
}

func renderJSON(resp http.ResponseWriter, status int, obj interface{}) {
	var result []byte
	var err error

	result, err = json.Marshal(obj)
	if err != nil {
		renderInternalError(resp, err)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	if _, err = io.Copy(resp, bytes.NewBuffer(result)); err != nil {
		renderInternalError(resp, err)
		return
	}
}
