package main

import (
		"context"
    "log"
		"github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/sqs"
		"fmt"
)

func main() {
	ctx := context.Background()
	queue_name := "urutan"
	endpoint := "http://localhost:4566"
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
			log.Fatal(err)
	}

	fmt.Println("queueName: %s endpoint: %s \n ", queue_name, endpoint)
	// Create an Amazon SQS service client
	client := sqs.NewFromConfig(cfg, func (o *sqs.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})

	queue, err := client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String(queue_name),
	})
	if err != nil {
			log.Fatal(err)
	}

	_, err = client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String("Hello SQS!"),
		QueueUrl:	queue.QueueUrl,
	})
	if err != nil {
			log.Fatal(err)
	}
}