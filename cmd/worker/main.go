package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/RobsonFeitosa/go-driver/internal/bucket"
	"github.com/RobsonFeitosa/go-driver/internal/queue"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func main() {
	qcfg := queue.RabbitMQConfig{
		URL:       "amqp://" + os.Getenv("RABBIT_URL"),
		TopicName: os.Getenv("RABBIT_TOPIC_NAME"),
		Timeout:   time.Second * 30,
	}

	fmt.Println("entrou1")

	qc, err := queue.New(queue.RabbitMQ, qcfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("entrou2")
	c := make(chan queue.QueueDto, 1)
	go qc.Consume(c)

	bcfg := bucket.AwsConfig{
		Config: &aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET"), ""),
		},
		BucketDownload: "aprenda-golang-drive-raw",
		BucketUpload:   "aprenda-golang-drive-gzip",
	}

	fmt.Println("entrou3")

	b, err := bucket.New(bucket.AwsProvider, bcfg)
	if err != nil {
		panic(err)
	}

	log.Println("waiting for messages")
	for msg := range c {
		dst := fmt.Sprintf("%d_%s", msg.ID, msg.Filename)

		log.Printf("Start working on %s\n", msg.Filename)

		err := b.Download(msg.Path, dst)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}
		fmt.Println("entrou4", dst)

		file, err := os.Open(dst)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		fmt.Println("entrou5")
		body, err := io.ReadAll(file)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		fmt.Println("entrou6")
		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		_, err = zw.Write(body)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		fmt.Println("entrou7")
		if err := zw.Close(); err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}
		fmt.Println("entrou8")

		zr, err := gzip.NewReader(&buf)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		fmt.Println("entrou9", zr, msg.Path)

		err = b.Upload(zr, msg.Path)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		fmt.Println("entrou10")

		err = os.Remove(dst)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		log.Printf("%s was proccesed with success!\n", msg.Filename)
	}
}
