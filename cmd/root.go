package cmd

import (
	"fmt"
	"os"

	"github.com/babyfaceeasy/crims/internal/seeds"
	"github.com/babyfaceeasy/crims/internal/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crims", // Root command name
	Short: "CLI for CRIMS",
	Long:  `Cloud Resource Information Management System.`,
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with initial data",
	Run: func(cmd *cobra.Command, args []string) {
		err := server.LoadEnv()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		seeds.Execute(args...) // Execute the seed logic, passing any remaining args
	},
}

func Execute() {
	rootCmd.AddCommand(seedCmd)

	// Parse the command-line arguments
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
func Execute() error {
	return rootCmd.Execute()
}
*/
