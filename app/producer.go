package app

import "github.com/nsqio/go-nsq"

func NewProducer() *nsq.Producer {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		panic(err)
	}
	return producer
}
