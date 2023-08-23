// commands/ts.go
package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/louiseschmidtgen/KVDB/database"

	"github.com/spf13/cobra"
)

func NewTimestampCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ts",
		Short: "Get the timestamp of a key",
		Run:   Timestamp,
	}
	return cmd
}

func Timestamp(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: kvdb ts <key>")
		os.Exit(1)
	}

	key := args[0]

	db, err := database.InitKeyValueDB("database/kvdb.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	timestamp := db.Timestamp(key)
	if timestamp == nil {
		fmt.Println("Key not found")
	} else {
		fmt.Printf("First set at: %s\n", timestamp[0].Format(time.RFC3339))
		fmt.Printf("Last set at: %s\n", timestamp[len(timestamp)-1].Format(time.RFC3339))
	}
}
