package main

import (
	"SparkGuardBackend/controllers"
	"log"
)

func main() {
	r := controllers.SetupRouter()

	log.Println("Server started")
	r.Run(":8080")
}
