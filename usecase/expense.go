package usecase

import (
	"context"
	"fmt"

	"github.com/suda7kabo/household-account-book/domain/object/expense"
	"github.com/suda7kabo/household-account-book/domain/repository"
)

type ExpenseUseCase interface {
	Create(ctx context.Context, name string) error
}

type useCase struct {
	r repository.Expense
}

func NewExpenseUseCase(r repository.Expense) ExpenseUseCase {
	return &useCase{
		r: r,
	}
}

func (u useCase) Create(ctx context.Context, name string) error {
	e, err := expense.NewExpense(name)
	if err != nil {
		return fmt.Errorf("failed to generate expense: %w", err)
	}

	if err := u.r.Create(ctx, e); err != nil {
		return fmt.Errorf("failed to create expense:%w", err)
	}
	return nil
}
