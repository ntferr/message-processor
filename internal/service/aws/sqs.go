package aws

import (
	"encoding/json"
	"log"
	"message-processor/internal/models"
	"message-processor/internal/settings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

func CreateQueueClient() sqsiface.SQSAPI {
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sqs.New(awsSession)
}

func SendMessage(user models.User) {
	message, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("ain")
	}

	svc := CreateQueueClient()

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Author": {
				DataType:    aws.String("String"),
				StringValue: aws.String("Nathan"),
			},
		},
		MessageBody: aws.String(string(message)),
		QueueUrl:    aws.String(settings.GetSettings().AWS.URL),
	})

	if err != nil {
		log.Fatalf("Error on message sending %v", err)
	}
}
