//go:generate mockgen -destination mock_expense.go -mock_names Expense=MockExpense -package repository "github.com/suda7kabo/household-account-book/domain/repository" Expense

package repository
