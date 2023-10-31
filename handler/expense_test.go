package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"github.com/suda7kabo/household-account-book/util/logs"

	"github.com/golang/mock/gomock"
	"github.com/suda7kabo/household-account-book/test/mock/usecase"
	expenseUseCase "github.com/suda7kabo/household-account-book/usecase"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestExpenseHandler_CreateExpense(t *testing.T) {
	// ------setup------
	e := echo.New()
	ctx := context.Background()
	useCaseMock := usecase.NewMockExpenseUseCase(gomock.NewController(t))
	logger, err := logs.NewLogger()
	assert.NoError(t, err)
	h := NewExpenseHandler(useCaseMock, logger)

	name := "娯楽費"
	id := "66babf7c-0748-4fa8-b5f0-a7b571c5dd19"
	// ------setup------

	type resp struct {
		statusCode int
		body       func() string
	}

	tests := []struct {
		name     string
		mock     func()
		request  func() []byte
		response resp
	}{
		{
			name: "正常時、正常系のレスポンスを返却すること",
			mock: func() {
				useCaseMock.EXPECT().Create(ctx, gomock.Eq(name)).Return(
					&expenseUseCase.ExpenseDTO{
						ID:   id,
						Name: name,
					}, nil)
			},
			request: func() []byte {
				reqBody, err := json.Marshal(&ReqBodyExpense{
					Name: name,
				})
				assert.NoError(t, err)
				return reqBody
			},
			response: resp{
				statusCode: http.StatusCreated,
				body: func() string {
					s, err := json.Marshal(&Expense{
						ID:   id,
						Name: name,
					})
					assert.NoError(t, err)
					return string(s)
				},
			},
		},
		{
			name: "Bind()でエラー発生時、BadRequestエラーを返却すること",
			mock: func() {},
			request: func() []byte {
				return []byte("invalidReq")
			},
			response: resp{
				statusCode: http.StatusBadRequest,
				body: func() string {
					s, err := json.Marshal(&HTTPError{
						Code:    001,
						Message: "パラメータのバインドに失敗しました。",
					})
					assert.NoError(t, err)
					return string(s)
				},
			},
		},
		{
			name: "usecase.Create()でエラー発生時、InternalServerエラーを返却すること",
			mock: func() {
				useCaseMock.EXPECT().Create(ctx, gomock.Eq(name)).Return(
					nil, errors.New("some error happened"))
			},
			request: func() []byte {
				reqBody, err := json.Marshal(&ReqBodyExpense{
					Name: name,
				})
				assert.NoError(t, err)
				return reqBody
			},
			response: resp{
				statusCode: http.StatusInternalServerError,
				body: func() string {
					s, err := json.Marshal(&HTTPError{
						Code:    002,
						Message: "費目の登録に失敗しました。",
					})
					assert.NoError(t, err)
					return string(s)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			httpReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(tt.request()))
			httpReq.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(httpReq.WithContext(ctx), rec)

			err := h.CreateExpense(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.response.statusCode, rec.Code)
			assert.JSONEq(t, tt.response.body(), rec.Body.String())
		})
	}
}
