package nsq

import (
	nsq "github.com/nsqio/go-nsq"
)

type Repository struct {
	nsq *nsq.Producer
}

func New(nsq *nsq.Producer) *Repository {
	return &Repository{nsq}
}
