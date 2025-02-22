package main

import (
	"SparkGuardBackend/cmd/orchestrator/controllers"
	"SparkGuardBackend/pkg/s3storage"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := s3storage.Connect(os.Getenv("S3_ENDPOINT"), os.Getenv("S3_REGION"), os.Getenv("S3_BUCKET")); err != nil {
		panic(err)
	}

	r := controllers.SetupRouter()

	log.Println("Server started")
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error:", err)
	}
}
