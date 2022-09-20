package main

import (
	"crud/internal/config"
	"crud/internal/factory"
	m "crud/internal/middleware"
	"crud/internal/routes"
)

func main() {
	config.InitDB()
	f := factory.NewFactory()
	e := routes.New(f)
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8082"))

}
