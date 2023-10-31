package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/suda7kabo/household-account-book/domain/object/expense"
	"github.com/suda7kabo/household-account-book/test/mock/repository"
)

func TestExpenseUseCase_Create(t *testing.T) {
	// ------setup------
	ctx := context.Background()
	repositoryMock := repository.NewMockExpense(gomock.NewController(t))
	uc := NewExpenseUseCase(repositoryMock)
	now := time.Now()
	id := "66babf7c-0748-4fa8-b5f0-a7b571c5dd19"
	name := "娯楽費"
	// ------setup------

	type input struct {
		name string
	}

	type want struct {
		res *ExpenseDTO
		err error
	}

	tests := []struct {
		name  string
		input input
		mock  func()
		want  want
	}{
		{
			name: "正常時、nilが返ること",
			input: input{
				name: name,
			},
			mock: func() {
				repositoryMock.EXPECT().Create(ctx, gomock.Any()).Return(&expense.Expense{
					ID:        expense.ID(id),
					Name:      expense.Name(name),
					CreatedAt: now,
					UpdatedAt: now,
				}, nil)
			},
			want: want{
				res: &ExpenseDTO{
					ID:   id,
					Name: name,
				},
				err: nil,
			},
		},
		{
			name: "引数のnameの値が不正な場合、エラーが返ること",
			input: input{
				name: "invalidNameOverLimitLength",
			},
			mock: func() {},
			want: want{
				res: nil,
				err: errors.New("failed to generate expense: failed to generate name: expense name must be 15 characters or less"),
			},
		},
		{
			name: "repository.Create()でエラー発生時、エラーが返ること",
			input: input{
				name: name,
			},
			mock: func() {
				repositoryMock.EXPECT().Create(ctx, gomock.Any()).Return(nil, errors.New("some error happened"))
			},
			want: want{
				res: nil,
				err: errors.New("failed to create expense: some error happened"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := uc.Create(ctx, tt.input.name)
			if tt.want.err == nil {
				assert.Equal(t, tt.want.res, got)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.want.err.Error())
				assert.Nil(t, tt.want.res)
			}
		})
	}

}
