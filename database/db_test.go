package database_test

import (
	"os"
	"testing"
	"time"

	"github.com/louiseschmidtgen/KVDB/database"
	"github.com/stretchr/testify/require"
)

func TestKeyValueDB(t *testing.T) {
	t.Parallel() // Indicate that this test can run in parallel
	// Test InitKeyValueDB
	db, err := database.InitKeyValueDB("test.db")
	require.NoError(t, err)

	// Close and remove the database when the function returns
	defer os.Remove("test.db")
	defer db.Close()

	// Test Set and Get
	err = db.Set("key1", "value1")
	require.NoError(t, err)

	value, err := db.Get("key1")

	require.NoError(t, err)
	require.Equal(t, "value1", value) // Confirm that the value is correct

	// Test Delete
	found := db.Delete("key1")
	require.True(t, found)

	// Confirm that the key-vslue pair has been deleted
	value, err = db.Get("key1")
	require.ErrorContains(t, err, "not found")
	require.Equal(t, "", value)

	// Test Timestamp
	// Create a key-value pair
	db.Set("key2", "value2")
	time.Sleep(time.Millisecond) // Sleep to ensure different timestamps
	// update the key with a new value which should update the timestamp
	db.Set("key2", "value2_updated")
	timestamp, err := db.Timestamp("key2")

	require.NoError(t, err)
	require.Len(t, timestamp, 2)                       // Confirm that there are two timestamps
	require.True(t, timestamp[0].Before(timestamp[1])) // Confirm that the first timestamp is before the second
}

func TestKeyValueDB_NotFound(t *testing.T) {
	t.Parallel() // Indicate that this test can run in parallel

	db, err := database.InitKeyValueDB("test2.db")

	require.NoError(t, err)

	defer os.Remove("test2.db")
	defer db.Close()

	// Test Get with a non-existent key
	value, err := db.Get("nonexistent_key")
	require.ErrorContains(t, err, "not found")
	require.Equal(t, "", value)

	// Test Timestamp with a non-existent key
	timestamp, err := db.Timestamp("nonexistent_key")
	require.ErrorContains(t, err, "not found")
	require.Nil(t, timestamp)
}
