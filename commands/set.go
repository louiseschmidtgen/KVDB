package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/louiseschmidtgen/KVDB/database"

	"github.com/spf13/cobra"
)

func NewSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set a key-value pair",
		Run:   SetCmdWrapper,
	}

	// Add an optional filename flag with default value
	cmd.Flags().String("filename", "database/kvdb.db", "Database filename")
	return cmd
}

func SetCmdWrapper(cmd *cobra.Command, args []string) {
	// Since you can not pass an error back to a cobra command from a function
	// but I would still like to do error handling so I have added a Wrapper function
	err := Set(cmd, args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
func Set(cmd *cobra.Command, args []string) error {
	// Check command line arguments
	if len(args) != 2 {
		return fmt.Errorf("Usage: kvdb set <key> <value>")
	}

	key := args[0]
	value := args[1]

	filename := cmd.Flag("database").Value.String()

	// Open the database
	db, err := database.InitKeyValueDB(filename)
	if err != nil {
		return fmt.Errorf("Error opening database: %v", err)
	}

	// Close the database when the function returns
	defer db.Close()

	// Set the value for the key
	err = db.Set(key, value)
	if err != nil {
		return fmt.Errorf("Error setting the value for %s: %v", value, err)
	}

	return nil
}
