// commands/del.go
package commands

import (
	"errors"
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
	if err := Delete(cmd, args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func Delete(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: kvdb del <key>")
	}

	key := args[0]
	filename := cmd.Flag("database").Value.String()

	// Open the database
	db, err := database.InitKeyValueDB(filename)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	// Close the database when the function returns
	defer db.Close()

	// Delete the key
	if found := db.Delete(key); found {
		log.Printf("Key '%s' deleted\n", key)
	} else {
		return errors.New("key not found")
	}

	log.Printf("Key '%s' deleted\n", key)

	return nil
}
