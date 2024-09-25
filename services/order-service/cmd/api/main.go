package main

import (
	"github.com/serenite11/market/services/order-service/internal/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(app.CreateApp()).Run()
}
