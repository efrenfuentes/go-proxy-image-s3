package s3

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	log "github.com/sirupsen/logrus"
)

// key=uploads/12960da9f702fe19dbf11d2f8accb8094456b099.jpeg
func DownloadImage(key string, log *log.Logger) (*string, error) {
	filename := strings.ReplaceAll(key, "/", "_")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Errorf("Couldn't load SDK config, %v", err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	awsBucket := os.Getenv("AWS_BUCKET")

	output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(awsBucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Errorf("Couldn't get object %v. Here's why: %v\n", key, err)
		return nil, err
	}

	defer output.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		log.Errorf("Couldn't create file %v. Here's why: %v\n", filename, err)
		return nil, err
	}

	defer file.Close()

	body, err := io.ReadAll(output.Body)
	if err != nil {
		log.Errorf("Couldn't read object body from %v. Here's why: %v\n", key, err)
		return nil, err
	}
	_, err = file.Write(body)
	if err != nil {
		log.Errorf("Couldn't write to file %v. Here's why: %v\n", filename, err)
		return nil, err
	}

	return &filename, nil

}
