package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var arn_topic string = "arn:aws:sns:us-east-1:321829329018:Snakemake"

func main() {
	if len(os.Args) < 2 {
		log.Println("[Please specify a message]")
		return
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("us-east-1")},
		Profile: "gonotify",
	})
	if err != nil {
		log.Printf("[error] %s", err)
		os.Exit(1)
	}

	svc := sns.New(sess)

	var msg string = os.Args[1]
	var subj string = "Snakemake info"

	publish_input := sns.PublishInput{
		Message:  &msg,
		Subject:  &subj,
		TopicArn: &arn_topic}

	resp, err := svc.Publish(&publish_input)
	if err != nil {
		log.Printf("[error] %s", err)
		os.Exit(1)
	}

	log.Printf("[success] %s", resp)

}
