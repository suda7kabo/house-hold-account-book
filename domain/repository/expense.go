package repository

import (
	"context"

	"github.com/suda7kabo/household-account-book/domain/object/expense"
)

type Expense interface {
	Create(ctx context.Context, expense *expense.Expense) error
}
