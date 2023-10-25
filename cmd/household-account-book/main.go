package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/suda7kabo/household-account-book/handler"
)

func main() {
	e := echo.New()
	expenseHandler := handler.NewExpenseHandler()

	e.POST("/expenses", expenseHandler.CreateExpense)
	go func() {
		if err := e.Start(":1323"); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		panic(fmt.Errorf("failed to graceful shutdown:%w", err))
	}
	log.Println("Server shutdown")
}
