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
	now := time.Now()
	seconds := now.Add(30 * time.Second)

	qcfg := queue.RabbitMQConfig{
		URL:       os.Getenv("RABBIT_URL"),
		TopicName: os.Getenv("RABBIT_TOP_NAME"),
		Timeout:   seconds,
	}

	// create new queue
	qc, err := queue.New(queue.RabbitMQ, qcfg)
	if err != nil {
		panic(err)
	}

	// create channel to consume message
	c := make(chan queue.QueueDto)
	qc.Consume(c)

	// bucket config
	bcfg := bucket.AwsConfig{
		Config: &aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET"), ""),
		},
		BucketDownload: "golang-drive-raw",
		BucketUpload:   "golang-drive-gzip",
	}

	// create new bucket session
	b, err := bucket.New(bucket.AwsProvider, bcfg)
	if err != nil {
		panic(err)
	}

	for msg := range c {
		src := fmt.Sprintf("%s/%s", msg.Path, msg.Filename)
		dst := fmt.Sprintf("%d_%s", msg.ID, msg.Filename)
		file, err := b.Download(src, dst)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		body, err := io.ReadAll(file)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}

		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		_, err = zw.Write(body)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		if err := zw.Close(); err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		zr, err := gzip.NewReader(&buf)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		err = b.Upload(zr, src)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		err = os.Remove(dst)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}
	}
}
