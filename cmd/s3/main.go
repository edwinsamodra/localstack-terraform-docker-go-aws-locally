package main

import (
    "context"
    "log"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
		"io"
		"fmt"
)

func main() {
    // Load the Shared AWS Configuration (~/.aws/config)
		ctx := context.Background()
		bucketName := "ember-tess"
    cfg, err := config.LoadDefaultConfig(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // Create an Amazon S3 service client
    client := s3.NewFromConfig(cfg, func (o *s3.Options) {
			o.BaseEndpoint = aws.String("http://s3.localhost.localstack.cloud:4566")
	
		})

    // Get the first page of results for ListObjectsV2 for a bucket
    output, err := client.GetObject(ctx, &s3.GetObjectInput{
        Bucket: aws.String(bucketName),
				Key: aws.String("go.mod"),
    })
    if err != nil {
        log.Fatal(err)
    }

		b, err := io.ReadAll(output.Body)
		if err != nil {
        log.Fatal(err)
    }
		fmt.Println(string(b))
}
