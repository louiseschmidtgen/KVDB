// commands/ts.go
package commands

import (
	"errors"
	"fmt"
	"log"
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
	if err := Timestamp(cmd, args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func Timestamp(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: kvdb ts <key>")
	}

	key := args[0]

	filename := cmd.Flag("database").Value.String()

	db, err := database.InitKeyValueDB(filename)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}
	defer db.Close()

	timestamp, err := db.Timestamp(key)
	if err != nil {
		return err
	}

	log.Printf("First set at: %s\n", timestamp[0].Format(time.RFC3339))
	log.Printf("Last set at: %s\n", timestamp[len(timestamp)-1].Format(time.RFC3339))

	return nil
}
