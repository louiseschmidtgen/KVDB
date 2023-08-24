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

const (
	key              = "greeting"
	value            = "hi"
	filenameFlagName = "database"
)

func TestMain(m *testing.M) {
	// Set up the test database once
	db, err := database.InitKeyValueDB("test3.db")
	if err != nil {
		log.Fatalf("error initializing test database: %v", err)
	}

	// Run all the tests
	code := m.Run()

	// Clean up after all tests are done
	db.Close()
	os.Remove("test3.db")

	// Exit with the test result code
	os.Exit(code)
}

func TestCommandsWithFlag(t *testing.T) {
	t.Parallel() // Run this test in parallel
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

	require.Contains(t, stdoutBuffer.String(), "First set at:")
	require.Contains(t, stdoutBuffer.String(), "Last set at:")
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
	t.Parallel() // Run this test in parallel
	t.Run("SetCommand", func(t *testing.T) {
		t.Parallel() // Run this subtest in parallel
		testSetCommandCreation(t)
	})

	t.Run("TSCommand", func(t *testing.T) {
		t.Parallel() // Run this subtest in parallel
		testTSCommandCreation(t)
	})

	t.Run("GetCommand", func(t *testing.T) {
		t.Parallel() // Run this subtest in parallel
		testGetCommandCreation(t)
	})

	t.Run("DelCommand", func(t *testing.T) {
		t.Parallel() // Run this subtest in parallel
		testDelCommandCreation(t)
	})
}

func testSetCommandCreation(t *testing.T) {
	t.Helper() // Add this line to tell the test framework that this is a helper function

	cmd := commands.NewSetCommand()
	require.Equal(t, "set", cmd.Use)
	require.Equal(t, "Set a key-value pair", cmd.Short)
}

func testTSCommandCreation(t *testing.T) {
	t.Helper()

	cmd := commands.NewTimestampCommand()
	require.Equal(t, "ts", cmd.Use)
	require.Equal(t, "Get the timestamp of a key", cmd.Short)
}

func testGetCommandCreation(t *testing.T) {
	t.Helper()

	cmd := commands.NewGetCommand()
	require.Equal(t, "get", cmd.Use)
	require.Equal(t, "Get a key-value pair", cmd.Short)
}

func testDelCommandCreation(t *testing.T) {
	t.Helper()

	cmd := commands.NewDeleteCommand()
	require.Equal(t, "del", cmd.Use)
	require.Equal(t, "Delete a key and its associated value", cmd.Short)
}
