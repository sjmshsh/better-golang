package main

import (
	"os"
)

const (
	Topic     = "xxxxx"
	Endpoint  = "xxxxx"
	AccessKey = "xxxxx"
	SecretKey = "xxxxx"
)

func main() {
	os.Setenv("mq.consoleAppender.enabled", "true")

}
