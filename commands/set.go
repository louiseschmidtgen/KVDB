// commands/set.go
package commands

import (
	"fmt"
	"os"

	"github.com/louiseschmidtgen/KVDB/database"

	"github.com/spf13/cobra"
)

func NewSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set a key-value pair",
		Run:   Set,
	}
	return cmd
}

func Set(cmd *cobra.Command, args []string) {
	fmt.Printf("args: %v\n", args)
	if len(args) != 2 {
		fmt.Println("Usage: kvdb set <key> <value>")
		os.Exit(1)
	}

	key := args[0]
	value := args[1]

	db, err := database.InitKeyValueDB("database/kvdb.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	db.Set(key, value)
	fmt.Printf("Key '%s' set to '%s'\n", key, value)
}
