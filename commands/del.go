// commands/del.go
package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/louiseschmidtgen/KVDB/database"

	"github.com/spf13/cobra"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del",
		Short: "Delete a key and its associated value",
		Run:   GetDelWrapper,
	}
	// Add an optional database flag with default value
	cmd.Flags().String("database", "database/kvdb.db", "Database filename")
	return cmd
}

func GetDelWrapper(cmd *cobra.Command, args []string) {
	// Since you can not pass an error back to a cobra command from a function
	// but I would still like to do error handling so I have added a Wrapper function
	err := Get(cmd, args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
func Delete(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: kvdb del <key>")
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

	// Delete the key
	found := db.Delete(key)

	if found {
		log.Printf("Key '%s' deleted\n", key)
	} else {
		return fmt.Errorf("Key not found")
	}

	fmt.Printf("Key '%s' deleted\n", key)
	return nil
}
