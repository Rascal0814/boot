package main

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"boot"
	"boot/cmd/project"
)

func main() {
	root := &cobra.Command{
		Use:           "boot",
		Short:         "An elegant toolkit for Golang microservice",
		Version:       boot.Version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.AddCommand(project.CommandNew())

	if err := root.ExecuteContext(context.Background()); err != nil {
		log.Fatal(err)
	}
}
