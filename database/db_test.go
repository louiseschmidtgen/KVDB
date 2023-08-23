package database

import (
	"os"
	"testing"
	"time"
)

func TestKeyValueDB(t *testing.T) {
	db, err := InitKeyValueDB("test.db")
	if err != nil {
		t.Fatalf("Failed to create KeyValueDB: %v", err)
	}
	defer os.Remove("test.db")
	defer db.Close()

	// Test Set and Get
	db.Set("key1", "value1")
	value, timestamp := db.Get("key1")
	if value != "value1" {
		t.Errorf("Expected value 'value1', got '%s'", value)
	}
	if len(timestamp) != 1 {
		t.Errorf("Expected timestamp length 1, got %d", len(timestamp))
	}

	// Test Delete
	db.Delete("key1")
	_, timestamp = db.Get("key1")
	if len(timestamp) != 0 {
		t.Errorf("Expected timestamp length 0 after delete, got %d", len(timestamp))
	}

	// Test Timestamp
	db.Set("key2", "value2")
	time.Sleep(time.Millisecond) // Sleep to ensure different timestamps
	db.Set("key2", "value2_updated")
	timestamp = db.Timestamp("key2")
	if len(timestamp) != 2 {
		t.Errorf("Expected timestamp length 2, got %d", len(timestamp))
	}
	if timestamp[0].After(timestamp[1]) {
		t.Errorf("Expected timestamp[0] before timestamp[1]")
	}
}

func TestKeyValueDB_NotFound(t *testing.T) {
	db, err := InitKeyValueDB("test.db")
	if err != nil {
		t.Fatalf("Failed to create KeyValueDB: %v", err)
	}
	defer os.Remove("test.db")
	defer db.Close()

	value, timestamp := db.Get("nonexistent_key")
	if value != "" {
		t.Errorf("Expected value '', got '%s'", value)
	}
	if timestamp != nil {
		t.Errorf("Expected timestamp nil, got %v", timestamp)
	}

	timestamp = db.Timestamp("nonexistent_key")
	if timestamp != nil {
		t.Errorf("Expected timestamp nil, got %v", timestamp)
	}
}
