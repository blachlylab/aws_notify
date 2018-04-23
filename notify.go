package main

import (
	"flag"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	var topic = flag.String("topic", "arn:aws:sns:us-east-1:321829329018:Snakemake", "Specify a complete Topic ARN or give only the final identifier for auto-lookup")
	flag.Parse()

	if flag.NArg() < 1 {
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

	var msg string = flag.Arg(0)
	var subj string = "Snakemake info"

	publish_input := sns.PublishInput{
		Message:  &msg,
		Subject:  &subj,
		TopicArn: topic} // no & operator because topic is a ptr

	resp, err := svc.Publish(&publish_input)
	if err != nil {
		log.Printf("[error] %s", err)
		os.Exit(1)
	}

	log.Printf("[success] %s", resp)

}
