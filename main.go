package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jgroeneveld/bookie/db"
	"github.com/jgroeneveld/bookie/handlers"
)

func main() {
	dbConnection := db.OpenDb()
	defer dbConnection.Close()

	router := mux.NewRouter()
	expensesHandler := handlers.ExpensesHandler{DB: dbConnection}

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Path("/expenses").Methods("GET").HandlerFunc(expensesHandler.GetExpenses)
	apiRouter.Path("/expenses").Methods("POST").HandlerFunc(expensesHandler.CreateExpense)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", router)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
