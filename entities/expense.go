package entities

import "time"

type Expense struct {
	ID        int
	Category  Category
	Amount    Money
	CreatedAt time.Time
	SpentAt   time.Time
	User      User
}
