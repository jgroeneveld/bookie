package entities

type ExpensesReport struct {
	AmountByUsers  UserMoneyMap
	MonthlyReports []MonthlyReport
}

type MonthlyReport struct {
	Month            Month
	TotalAmount      Money
	AmountByUsers    UserMoneyMap
	AmountByCategory CategoryMoneyMap
}
