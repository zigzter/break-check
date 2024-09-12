package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/zigzter/break-check/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Termination signal received. Cancelling...")
		cancel()
	}()

	if err := cmd.Execute(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
