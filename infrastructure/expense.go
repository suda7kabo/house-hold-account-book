package infrastructure

import (
	"context"
	"fmt"

	"github.com/suda7kabo/household-account-book/domain/object/expense"
	"github.com/suda7kabo/household-account-book/domain/repository"
)

type expenseRepositoryWrapper struct {
	db *DB
}

func NewExpenseRepository(db *DB) repository.Expense {
	return &expenseRepositoryWrapper{
		db: db,
	}
}

func (w expenseRepositoryWrapper) Create(ctx context.Context, expense *expense.Expense) (*expense.Expense, error) {
	q := `
	INSERT INTO expenses (
		id,
		name,
		created_at,
		updated_at
	) VALUES(?, ?, ?, ?)`

	if _, err := w.db.Write.ExecContext(ctx, q, expense.ID, expense.Name.String(), expense.CreatedAt, expense.UpdatedAt); err != nil {
		return nil, fmt.Errorf("failed to ExecContext: %w", err)
	}
	return expense, nil
}
