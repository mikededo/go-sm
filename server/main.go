package main

import rest "github.com/mddg/go-sm/server/infrastructure/rest/router"

func main() {
	err := rest.NewRestRouter().Run(":3000")
	if err != nil {
		panic(err.Error())
	}
}
