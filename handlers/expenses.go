package handlers

import (
	"time"

	"github.com/jgroeneveld/bookie/entities"
	"github.com/martini-contrib/render"
)

func GetExpenses(render render.Render) {
	expenses := []entities.Expense{
		{Category: "Edeka", Amount: 18.23, CreatedAt: dateFromString("2014-05-12"), User: "Hilke"},
		{Category: "Lidl", Amount: 5.19, CreatedAt: dateFromString("2014-05-17"), User: "Hilke"},
		{Category: "Edeka", Amount: 14.85, CreatedAt: dateFromString("2014-06-02"), User: "Jaap"},
		{Category: "Edeka", Amount: 22.42, CreatedAt: dateFromString("2014-06-07"), User: "Hilke"},
	}
	render.JSON(200, expenses)
}

func CreateExpense() string {
	return "Alles ut junge"
}

func dateFromString(s string) time.Time {
	layout := "2006-01-02"

	d, _ := time.Parse(layout, s)

	return d
}

func ExpensesReport(render render.Render) {
	report := entities.ExpensesReport{
		MonthlyReports: []entities.MonthlyReport{
			{
				Month:       "2014-05",
				TotalAmount: 22.23,
				AmountByUsers: entities.UserMoneyMap{
					"Jaap":  12.23,
					"Hilke": 10.00,
				},
				AmountByCategory: entities.CategoryMoneyMap{
					"Edeka": 12.23,
					"Lidl":  10.00,
				},
			},
		},
	}
	render.JSON(200, report)
}
