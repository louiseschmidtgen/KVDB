// commands/del.go
package commands

import (
	"fmt"
	"os"

	"github.com/louiseschmidtgen/KVDB/database"

	"github.com/spf13/cobra"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del",
		Short: "Delete a key and its associated value",
		Run:   Delete,
	}
	return cmd
}

func Delete(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: kvdb del <key>")
		os.Exit(1)
	}

	key := args[0]

	db, err := database.InitKeyValueDB("database/kvdb.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	if db.Delete(key) {
		fmt.Printf("Key '%s' deleted\n", key)
	} else {
		fmt.Println("Key not found")
	}
}
