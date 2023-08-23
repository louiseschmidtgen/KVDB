package main

import (
	"log"
	"os"

	"github.com/louiseschmidtgen/KVDB/commands"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "kvdb"}

	rootCmd.AddCommand(commands.NewSetCommand())
	rootCmd.AddCommand(commands.NewGetCommand())
	rootCmd.AddCommand(commands.NewDeleteCommand())
	rootCmd.AddCommand(commands.NewTimestampCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
