/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/chemi123/ldocker/cmd/subcommands"
	"github.com/chemi123/ldocker/pkg/factory"
	"github.com/spf13/cobra"
)

const (
	shortRootDesc = "A lightweight Docker management tool, focusing on simplifying container and image operations."

	longRootDesc = `ldocker is a lightweight command-line tool designed to streamline the management of Docker containers and images. With a focus on ease of use and simplicity, ldocker provides users with essential functionalities such as starting and stopping containers, listing and removing images, checking container statuses, and viewing logs. It aims to offer a more intuitive and straightforward approach to Docker management, making it an ideal choice for developers and system administrators who need to perform common Docker operations without the complexity of the full Docker CLI.
    
    Features include:
    - Starting and stopping containers: Quickly start up or shut down your Docker containers with simple commands.
    - Managing images: Easily list all available Docker images and remove the ones you no longer need.
    - Checking container statuses: Get a clear view of which containers are running and their current status.
    - Viewing logs: Access the logs of your containers to troubleshoot issues or monitor their operation.
    
    Designed with simplicity in mind, ldocker is the perfect tool for those who value efficiency and a clutter-free command-line experience in their Docker management tasks.`
)

func Execute() {
	ctx := context.Background()
	rootCmd := newRootCommand(ctx)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func newRootCommand(ctx context.Context) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ldocker",
		Short: shortRootDesc,
		Long:  longRootDesc,
	}

	// 仮に他のClientFactoryを注入するならここになる。
	clientFactory := factory.NewDockerdClientFactory()

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(subcommands.NewImagesSubcommand(ctx, clientFactory))

	return rootCmd
}
