package main

import (
	_ "github.com/joho/godotenv/autoload" // autoload env files
	rest "github.com/mddg/go-sm/server/infrastructure/rest/router"
)

func main() {
	err := rest.NewRestRouter().Run(":3000")
	if err != nil {
		panic(err.Error())
	}
}
