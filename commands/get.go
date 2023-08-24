// commands/get.go
package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/louiseschmidtgen/kvdb/database"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a key-value pair",
		Run:   GetCmdWrapper,
	}

	// Add an optional database flag with default value
	cmd.Flags().String("database", "database/kvdb.db", "Database filename")
	return cmd
}

func GetCmdWrapper(cmd *cobra.Command, args []string) {
	// Since you can not pass an error back to a cobra command from a function
	// but I would still like to do error handling so I have added a Wrapper function
	err := Get(cmd, args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func Get(cmd *cobra.Command, args []string) error {
	// Check command line arguments
	if len(args) != 1 {
		return fmt.Errorf("Usage: kvdb get <key>")
	}

	key := args[0]
	filename := cmd.Flag("database").Value.String()

	// Open the database
	db, err := database.InitKeyValueDB(filename)
	if err != nil {
		return fmt.Errorf("Error opening database: %v", err)
	}

	// Close the database when the function returns
	defer db.Close()

	// Get the value for the key
	value, err := db.Get(key)
	if err != nil {
		return err
	}

	// Print the value
	log.SetFlags(0)
	log.Printf("Value for '%s': %s\n", key, value)

	return nil
}
