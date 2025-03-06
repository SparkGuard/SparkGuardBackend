package main

import (
	"SparkGuardBackend/pkg/s3storage"
	"log"
	"os"
)

func main() {
	if err := s3storage.Connect(os.Getenv("S3_ENDPOINT"), os.Getenv("S3_REGION"), os.Getenv("S3_BUCKET")); err != nil {
		panic(err)
	}

	if err := serve(":666"); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
