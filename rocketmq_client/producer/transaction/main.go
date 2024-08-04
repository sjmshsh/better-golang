package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	Topic     = "xxxxxx"
	GroupName = "xxxxxx"
	Endpoint  = "xxxxxx"
	Region    = "xxxxxx"
	AccessKey = "xxxxxx"
	SecretKey = "xxxxxx"
)

func main() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	golang.ResetLogger()
	// new producer instance
	// new producer instance
	producer, err := golang.NewProducer(&golang.Config{
		Endpoint: Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		golang.WithTransactionChecker(&golang.TransactionChecker{
			Check: func(msg *golang.MessageView) golang.TransactionResolution {
				log.Printf("check transaction message: %v", msg)
				return golang.COMMIT
			},
		}),
		golang.WithTopics(Topic),
	)
	if err != nil {
		log.Fatal(err)
	}
	// start producer
	err = producer.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer producer.GracefulStop()
	for i := 0; i < 10; i++ {
		// new a message
		msg := &golang.Message{
			Topic: Topic,
			Body:  []byte("this is a message : " + strconv.Itoa(i)),
		}
		// set keys and tag
		msg.SetKeys("a", "b")
		msg.SetTag("ab")
		// send message in sync
		transaction := producer.BeginTransaction()
		resp, err := producer.SendWithTransaction(context.TODO(), msg, transaction)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(resp); i++ {
			fmt.Printf("%#v\n", resp[i])
		}
		// commit transaction message
		err = transaction.Commit()
		if err != nil {
			log.Fatal(err)
		}
		// wait a moment
		time.Sleep(time.Second * 1)
	}
}
