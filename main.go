package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("Custom CLI application")

	rootCmd := &cobra.Command{
		Use:   "greeter",
		Short: "A basic CLI example",
		Long:  "A basic CLI example using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Greeter!")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}
