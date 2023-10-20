package expense

import "fmt"

type Expense struct {
	ID   ID
	Name Name
}

func NewExpense(name string) (*Expense, error) {
	id, err := newID()
	if err != nil {
		return nil, fmt.Errorf("faild to generate new id: %w", err)
	}

	expenseName, err := newName(name)
	if err != nil {
		return nil, fmt.Errorf("faild to generate name: %w", err)
	}

	return &Expense{
		ID:   id,
		Name: expenseName,
	}, nil
}
