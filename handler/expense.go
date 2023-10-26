package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/suda7kabo/household-account-book/usecase"
)

type ExpenseHandler interface {
	CreateExpense(ctx echo.Context) error
}

type expenseHandlerWrapper struct {
	u usecase.ExpenseUseCase
}

func NewExpenseHandler(u usecase.ExpenseUseCase) ExpenseHandler {
	return &expenseHandlerWrapper{u: u}
}

func (w expenseHandlerWrapper) CreateExpense(ctx echo.Context) error {
	req := new(ReqBodyExpensee)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, HTTPError{
			Code:    001,
			Message: "パラメータのバインドに失敗しました。",
		})
	}
	result, err := w.u.Create(ctx.Request().Context(), req.Name)
	if err != nil {
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
