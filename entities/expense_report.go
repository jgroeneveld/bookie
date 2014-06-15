package entities

type ExpensesReport struct {
	MonthlyReports []MonthlyReport
}

type MonthlyReport struct {
	Month            Month
	TotalAmount      Money
	AmountByUsers    map[User]Money
	AmountByCategory map[Category]Money
}
