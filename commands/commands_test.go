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

var testDB *database.KeyValueDB
var key = "greeting"
var value = "hi"
var filenameFlagName = "database"

func TestMain(m *testing.M) {
	// Set up the test database once
	db, err := database.InitKeyValueDB("test3.db")
	if err != nil {
		log.Fatalf("Error initializing test database: %v", err)
	}
	testDB = db

	// Run all the tests
	code := m.Run()

	// Clean up after all tests are done
	testDB.Close()
	os.Remove("test3.db")

	// Exit with the test result code
	os.Exit(code)
}

func TestCommandsWithFlag(t *testing.T) {
	t.Run("SetCommand", testSetCommand)
	t.Run("TSCommand", testTSCommand)
	t.Run("GetCommand", testGetCommand)
	t.Run("DelCommand", testDelCommand)
}

func testSetCommand(t *testing.T) {
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
	cmd.Flags().String(filenameFlagName, "test3.db", "Database filename for testing")

	// Execute the SetCommand with the filename flag
	err := cmd.Execute()
	require.NoError(t, err)
}

func testTSCommand(t *testing.T) {
	// Use the shared testDB
	cmd := &cobra.Command{
		Use:   "ts",
		Short: "Get the timestamp of a key",
		Run: func(cmd *cobra.Command, args []string) {
			commands.Timestamp(cmd, args)
		},
	}

	// Set up test arguments for TSCommand
	k := []string{key}
	cmd.SetArgs(k)

	// Create a buffer to capture stdout
	var stdoutBuffer bytes.Buffer

	// Redirect stdout to the buffer
	log.SetOutput(&stdoutBuffer)

	cmd.Flags().String(filenameFlagName, "test3.db", "Database filename for testing")

	err := cmd.Execute()
	require.NoError(t, err)

	expectedOutput3 := "First set at:"
	expectedOutput4 := "Last set at:"
	require.Contains(t, stdoutBuffer.String(), expectedOutput3)
	require.Contains(t, stdoutBuffer.String(), expectedOutput4)

	// Initialize the GetCommand
	cmd = &cobra.Command{
		Use:   "set",
		Short: "Set a key-value pair",
		Run: func(cmd *cobra.Command, args []string) {
			commands.Get(cmd, args)
		},
	}

}

func testGetCommand(t *testing.T) {
	// Use the shared testDB
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a key-value pair",
		Run: func(cmd *cobra.Command, args []string) {
			commands.Get(cmd, args)
		},
	}

	// Create a buffer to capture stdout
	var stdoutBuffer bytes.Buffer

	// Redirect stdout to the buffer
	log.SetOutput(&stdoutBuffer)

	// Set up test arguments for GetCommand
	k := []string{key}
	cmd.SetArgs(k)

	// Set the filename flag for testing in GetCommand
	cmd.Flags().String(filenameFlagName, "test3.db", "Database filename for testing")

	// Execute the GetCommand with the filename flag
	err := cmd.Execute()
	require.NoError(t, err)

	// Check the content of the outputBuffer for our key and value
	expectedOutput := "Value for '" + key + "': " + value
	require.Contains(t, stdoutBuffer.String(), expectedOutput)
}

func testDelCommand(t *testing.T) {
	// Use the shared testDB
	cmd := &cobra.Command{
		Use:   "del",
		Short: "Delete a key and its associated value",
		Run: func(cmd *cobra.Command, args []string) {
			commands.Delete(cmd, args)
		},
	}

	// Create a buffer to capture stdout
	var stdoutBuffer bytes.Buffer

	// Redirect stdout to the buffer
	log.SetOutput(&stdoutBuffer)

	// Set up test arguments for DelCommand
	k := []string{key}
	cmd.SetArgs(k)

	// Set the filename flag for testing in DelCommand
	cmd.Flags().String(filenameFlagName, "test3.db", "Database filename for testing")

	// Execute the DelCommand with the filename flag
	err := cmd.Execute()
	require.NoError(t, err)

	// Check the content of the outputBuffer for our key
	expectedOutput := "Key '" + key + "' deleted"
	require.Contains(t, stdoutBuffer.String(), expectedOutput)
}

func TestCommandCreation(t *testing.T) {
	t.Run("SetCommand", testSetCommandCreation)
	t.Run("TSCommand", testTSCommandCreation)
	t.Run("GetCommand", testGetCommandCreation)
	t.Run("DelCommand", testDelCommandCreation)
}

func testSetCommandCreation(t *testing.T) {
	cmd := commands.NewSetCommand()
	require.Equal(t, "set", cmd.Use)
	require.Equal(t, "Set a key-value pair", cmd.Short)
}
func testTSCommandCreation(t *testing.T) {
	cmd := commands.NewTimestampCommand()
	require.Equal(t, "ts", cmd.Use)
	require.Equal(t, "Get the timestamp of a key", cmd.Short)
}
func testGetCommandCreation(t *testing.T) {
	cmd := commands.NewGetCommand()
	require.Equal(t, "get", cmd.Use)
	require.Equal(t, "Get a key-value pair", cmd.Short)
}
func testDelCommandCreation(t *testing.T) {
	cmd := commands.NewDeleteCommand()
	require.Equal(t, "del", cmd.Use)
	require.Equal(t, "Delete a key and its associated value", cmd.Short)
}
