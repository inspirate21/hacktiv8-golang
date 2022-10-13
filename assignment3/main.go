package main

import (
	"assignment3/routes"
	"log"
)

const PORT = ":3000"

func main() {
	router := routes.WeatherHttpHandler()
	log.Fatal(router.Run(PORT))
}
