package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jgroeneveld/bookie/db"
	"github.com/jgroeneveld/bookie/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("PORT missing")
	}

	dbConnection := db.OpenDb()
	defer dbConnection.Close()

	router := mux.NewRouter()
	expensesHandler := handlers.ExpensesHandler{DB: dbConnection}

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Path("/expenses").Methods("GET").HandlerFunc(expensesHandler.GetExpenses)
	apiRouter.Path("/expenses/report").Methods("GET").HandlerFunc(expensesHandler.GetExpensesReport)
	apiRouter.Path("/expenses").Methods("POST").HandlerFunc(expensesHandler.CreateExpense)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Printf("Serving on %s", port)
	http.ListenAndServe(":"+port, &LoggingHandler{Handler: router})
}

type LoggingHandler struct {
	Handler http.Handler
}

func (handler *LoggingHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s", req.Method, req.URL)
	handler.Handler.ServeHTTP(resp, req)
}
