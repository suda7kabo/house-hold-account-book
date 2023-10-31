package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/suda7kabo/household-account-book/usecase"
	"github.com/suda7kabo/household-account-book/util/logs"
)

type ExpenseHandler interface {
	CreateExpense(ctx echo.Context) error
	//ListExpense(ctx echo.Context) error
}

type expenseHandlerWrapper struct {
	l *logs.Logger
	u usecase.ExpenseUseCase
}

func NewExpenseHandler(u usecase.ExpenseUseCase, l *logs.Logger) ExpenseHandler {
	return &expenseHandlerWrapper{u: u, l: l}
}

func (w expenseHandlerWrapper) CreateExpense(ctx echo.Context) error {
	req := new(ReqBodyExpense)
	if err := ctx.Bind(req); err != nil {
		w.l.Error("failed to bind request", err)
		return ctx.JSON(http.StatusBadRequest, HTTPError{
			Code:    001,
			Message: "パラメータのバインドに失敗しました。",
		})
	}
	result, err := w.u.Create(ctx.Request().Context(), req.Name)
	if err != nil {
		w.l.Error("failed to create expense", err)
		return ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    002,
			Message: "費目の登録に失敗しました。",
		})
	}
	return ctx.JSON(http.StatusOK, Expense{
		ID:   result.ID,
		Name: req.Name,
	})
}
