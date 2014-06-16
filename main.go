package main

import (
	"github.com/go-martini/martini"
	"github.com/jgroeneveld/bookie/handlers"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	m.Use(martini.Static("public"))
	m.Use(render.Renderer())

	m.Group("/api/expenses", func(r martini.Router) {
		r.Get("", handlers.GetExpenses)
		r.Post("", handlers.CreateExpense)
		r.Get("/report", handlers.ExpensesReport)
	})

	m.Run()
}
