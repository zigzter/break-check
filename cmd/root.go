package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bc",
	Short: "Break Check collects all breaking changes in the given packages up to the latest version",
	Long:  "Break Check collects all breaking changes in the given packages up to the latest version",
}

func Execute() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Termination signal received. Cancelling...")
		cancel()
	}()

	return rootCmd.ExecuteContext(ctx)
}
