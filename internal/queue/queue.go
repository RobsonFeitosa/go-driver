package queue

import (
	"fmt"
	"log"
)

type QueueType int

const (
	RabbitMQ QueueType = iota
)

type QueueConnection interface {
	Publish([]byte) error
	Consume() error
}

type Queue struct {
	cfg any
	qc  QueueConnection
}

func New(qt QueueType, cfg any) *Queue {
	q := new(Queue)

	switch qt {
	case RabbitMQ:
		fmt.Println("Não implementado")
	default:
		log.Fatal("type not implemented")
	}

	return q
}
