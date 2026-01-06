package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

var ErrMissingAWSConfig = fmt.Errorf("missing AWS credentials or region")

// init() loads .env automatically
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}
}

func getPresignedURL(filename, contentType string) (string, error) {
	awsRegion := os.Getenv("AWS_REGION")
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	bucket := os.Getenv("S3_BUCKET")

	if awsRegion == "" || awsAccessKey == "" || awsSecretKey == "" || bucket == "" {
		log.Println("AWS credentials, region, or bucket not set")
		return "", ErrMissingAWSConfig
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, "")),
	)
	if err != nil {
		log.Println("Failed to load AWS config:", err)
		return "", err
	}

	client := s3.NewFromConfig(cfg)
	presigner := s3.NewPresignClient(client)

	req, err := presigner.PresignPutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(filename),
			ContentType: aws.String(contentType),
		},
		s3.WithPresignExpires(15*time.Minute),
	)
	if err != nil {
		log.Println("Failed to presign URL:", err)
		return "", err
	}

	log.Println("Presigned URL generated:", req.URL)
	return req.URL, nil
}
