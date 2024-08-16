package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myproject",
		Short: "A simple CLI project",
		Run: func(cmd *cobra.Command, args []string) {
			color.Info.Println("Hello from myproject!")
			if len(args) > 0 {
				fmt.Println("Arguments:", args)
			} else {
				fmt.Println("No arguments provided.")
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		color.Error.Println(err)
		os.Exit(1)
	}
}
