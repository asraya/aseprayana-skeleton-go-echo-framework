package main

import (
	"aseprayana-skeleton-go/configs/seeder"
	"aseprayana-skeleton-go/routes"
	"os"
)

func main() {

	dbEvent := os.Getenv("DBEVENT")
	if dbEvent == "seeder" {
		seeder.RunSeeder()
	}

	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
