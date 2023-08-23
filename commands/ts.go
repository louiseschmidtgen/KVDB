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
		Run:   TimestampWrapper,
	}

	// Add an optional database flag with default value
	cmd.Flags().String("database", "database/kvdb.db", "Database filename")
	return cmd
}

func TimestampWrapper(cmd *cobra.Command, args []string) {
	// Since you can not pass an error back to a cobra command from a function
	// but I would still like to do error handling so I have added a Wrapper function
	err := Timestamp(cmd, args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Timestamp(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: kvdb ts <key>")
	}

	key := args[0]

	filename := cmd.Flag("database").Value.String()
	db, err := database.InitKeyValueDB(filename)
	if err != nil {
		return fmt.Errorf("Error opening database:", err)

	}
	defer db.Close()

	timestamp, err := db.Timestamp(key)
	if err != nil {
		return err
	}

	fmt.Printf("First set at: %s\n", timestamp[0].Format(time.RFC3339))
	fmt.Printf("Last set at: %s\n", timestamp[len(timestamp)-1].Format(time.RFC3339))

	return nil
}
