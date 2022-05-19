package configs

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetSpacesClient() (*s3.S3, error) {
	key := os.Getenv("DO_ACCESS_KEY_ID")
	secret := os.Getenv("DO_SECRET_ACCESS_KEY")
	endpoint := os.Getenv("DO_ENDPOINT")
	

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint: aws.String(endpoint),
		Region: aws.String("us-east-1"),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}
	s3Client := s3.New(newSession)
	return s3Client, nil
}