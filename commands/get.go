// commands/get.go
package commands

import (
	"fmt"
	"os"

	"github.com/louiseschmidtgen/KVDB/database"

	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the value associated with a key",
		Run:   Get,
	}
	return cmd
}

func Get(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: kvdb get <key>")
		os.Exit(1)
	}

	key := args[0]

	db, err := database.InitKeyValueDB("database/kvdb.db")
	if err != nil {
		fmt.Println("Error opening the database:", err)
		os.Exit(1)
	}
	defer db.Close()

	value, _ := db.Get(key)
	if value == "" {
		fmt.Println("Key not found")
	} else {
		fmt.Printf("Value: %s\n", value)
	}
}
