package main

import (
	"SparkGuardBackend/cmd/rest/controllers"
	"SparkGuardBackend/pkg/s3storage"
	"log"
	"os"
)

func main() {
	if err := s3storage.Connect(os.Getenv("S3_ENDPOINT"), os.Getenv("S3_REGION"), os.Getenv("S3_BUCKET")); err != nil {
		panic(err)
	}

	r := controllers.SetupRouter()

	log.Println("Server started")
	r.Run(":8080")
}
