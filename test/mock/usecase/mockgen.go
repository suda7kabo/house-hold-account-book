//go:generate mockgen -destination mock_expense.go -mock_names Expense=MockExpense -package usecase "github.com/katsuharu/household-account-book/usecase" ExpenseUseCase

package usecase
