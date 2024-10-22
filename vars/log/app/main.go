package main

import (
	"context"

	"kwai/vars/log"
)

func main() {
	log.Init(log.NewOptions(log.WithLevel("debug")))
	log.Debug(context.Background(), "hehe")
}
