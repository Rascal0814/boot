package main

import (
	"context"
	"github.com/Rascal0814/boot"
	"github.com/Rascal0814/boot/cmd/boot/project"
	"log"

	"github.com/spf13/cobra"
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
