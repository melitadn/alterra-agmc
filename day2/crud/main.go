package main

import (
	"crud/config"
	"crud/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8082"))

}
