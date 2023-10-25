package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ExpenseHandler interface {
	CreateExpense(ctx echo.Context) error
}

type expenseHandlerWrapper struct {
}

func NewExpenseHandler() ExpenseHandler {
	return &expenseHandlerWrapper{}
}

func (w expenseHandlerWrapper) CreateExpense(ctx echo.Context) error {
	req := new(ReqBodyExpensee)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, HTTPError{
			Code:    001,
			Message: "パラメータのバインドに失敗しました。",
		})
	}
	return ctx.JSON(http.StatusOK, Expense{
		ID:   "expenseID",
		Name: req.Name,
	})
}
