package main

import (
	"context"

	"kwai/log"
)

func main() {
	log.Init(log.NewOptions(log.WithLevel("debug")))
	log.Debug(context.Background(), "hehe")
}
