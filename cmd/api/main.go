package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RobsonFeitosa/go-driver/internal/auth"
	"github.com/RobsonFeitosa/go-driver/internal/bucket"
	"github.com/RobsonFeitosa/go-driver/internal/files"
	"github.com/RobsonFeitosa/go-driver/internal/folders"
	"github.com/RobsonFeitosa/go-driver/internal/queue"
	"github.com/RobsonFeitosa/go-driver/internal/users"
	"github.com/RobsonFeitosa/go-driver/pkg/database"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/go-chi/chi"
)

func main() {
	db, b, qc := getSessions()

	r := chi.NewRouter()

	r.Post("/auth", auth.HandlerAuth(func(login, pass string) (auth.Authenticated, error) {
		return users.Authenticate(login, pass)
	}))

	files.SetRoutes(r, db, b, qc)
	folders.SetRoutes(r, db)
	users.SetRoutes(r, db)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), r)
}

func getSessions() (*sql.DB, *bucket.Bucket, *queue.Queue) {
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

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
		log.Fatal(err)
	}

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
		log.Fatal(err)
	}

	return db, b, qc
}
