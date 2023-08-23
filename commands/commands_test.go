package commands_test

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/louiseschmidtgen/KVDB/commands"
	"github.com/louiseschmidtgen/KVDB/database"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestSetGetDelCommandWithFlag(t *testing.T) {
	// Create and initialize a temporary test database
	db, err := database.InitKeyValueDB("test3.db")
	require.NoError(t, err)
	defer func() {
		db.Close()
		os.Remove("test3.db")
	}()

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set a key-value pair",
		Run: func(cmd *cobra.Command, args []string) {
			commands.Set(cmd, args)
		},
	}

	// Set up test arguments for SetCommand
	key := "greeting"
	value := "hi"
	kv := []string{key, value}
	cmd.SetArgs(kv)

	// Set the filename flag for testing
	filenameFlagName := "database"
	cmd.Flags().String(filenameFlagName, "test3.db", "Database filename for testing")

	// Execute the SetCommand with the filename flag
	err = cmd.Execute()
	require.NoError(t, err)

	// Initialize the GetCommand
	cmd = &cobra.Command{
		Use:   "set",
		Short: "Set a key-value pair",
		Run: func(cmd *cobra.Command, args []string) {
			commands.Get(cmd, args)
		},
	}

	// Set up test arguments for GetCommand
	k := []string{key}
	cmd.SetArgs(k)

	// Create a buffer to capture stdout
	var stdoutBuffer bytes.Buffer

	// Redirect stdout to the buffer
	log.SetOutput(&stdoutBuffer)

	// Set the filename flag for testing in GetCommand
	cmd.Flags().String(filenameFlagName, "test3.db", "Database filename for testing")

	// Execute the GetCommand with the filename flag
	err = cmd.Execute()
	require.NoError(t, err)

	// Check the content of the outputBuffer for our key and value
	expectedOutput := "Value for '" + key + "': " + value
	require.Contains(t, stdoutBuffer.String(), expectedOutput)

	cmd = &cobra.Command{
		Use:   "del",
		Short: "Delete a key and its associated value",
		Run: func(cmd *cobra.Command, args []string) {
			commands.Delete(cmd, args)
		},
	}

	// Set up test arguments for DeleteCommand
	cmd.SetArgs(k)

	cmd.Flags().String(filenameFlagName, "test3.db", "Database filename for testing")

	// Execute the DeleteCommand with the filename flag
	err = cmd.Execute()
	require.NoError(t, err)

	expectedOutput2 := "Key '" + key + "' deleted"
	require.Contains(t, stdoutBuffer.String(), expectedOutput2)
}
