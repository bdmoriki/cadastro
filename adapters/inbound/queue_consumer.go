package inbound

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

const (
	queueUrl = "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/fila-cadastro"
)

func QueueConsumer() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	sqsClient := sqs.NewFromConfig(cfg)

	queue, _ := sqsClient.GetQueueUrl(context.TODO(), &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueUrl),
	})

	messages, err := sqsClient.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:                queue.QueueUrl,
		AttributeNames:          nil,
		MaxNumberOfMessages:     10,
		MessageAttributeNames:   nil,
		ReceiveRequestAttemptId: nil,
		VisibilityTimeout:       0,
		WaitTimeSeconds:         0,
	})
	if err != nil {
		panic(err)
	}

	for _, message := range messages.Messages {
		log.Printf("the message body is %v", *message.Body)
	}
}
