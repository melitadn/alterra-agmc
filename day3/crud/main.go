package main

import (
	"crud/config"
	m "crud/middleware"
	"crud/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8082"))

}
