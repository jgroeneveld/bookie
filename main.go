package main

import (
	"github.com/go-martini/martini"
	"github.com/jgroeneveld/bookie/handlers"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	m.Use(martini.Static("assets"))
	m.Use(render.Renderer())

	m.Group("/expenses", func(r martini.Router) {
		r.Get("", handlers.GetExpenses)
		r.Get("/report", handlers.ExpensesReport)
	})

	m.Run()
}
