package outbound

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

const (
	topicArn = "arn:aws:sns:us-east-1:000000000000:cadastro"
)

func Publish() {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	client := sns.NewFromConfig(cfg, func(o *sns.Options) {
		o.BaseEndpoint = aws.String("http://localhost:4566")
	})

	publishInput := sns.PublishInput{
		TopicArn: aws.String(topicArn),
		Message:  aws.String("Teste")}

	_, err = client.Publish(context.TODO(), &publishInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Mensagem enviada com sucesso..")
}
