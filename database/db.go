package database

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"sync"
	"syscall"
	"time"
)

type KeyValueDB struct {
	sync.RWMutex
	data     map[string]DBEntry
	filename string
}

type DBEntry struct {
	Value     string
	Timestamp []time.Time
}

func InitKeyValueDB(filename string) (*KeyValueDB, error) {
	db := &KeyValueDB{
		data:     make(map[string]DBEntry),
		filename: filename,
	}

	if err := db.load(); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *KeyValueDB) load() error {
	// open the database file
	file, err := os.Open(db.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // KVDB does not exist yet
		}
		log.Printf("Error opening database file %s: %v\n", db.filename, err)
		return err
	}
	defer file.Close()

	// Lock the file for reading
	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_SH); err != nil {
		log.Printf("Error locking file %s for reading: %v\n", db.filename, err)
		return err
	}
	defer syscall.Flock(int(file.Fd()), syscall.LOCK_UN) // Unlock when done

	// decode the database file
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&db.data)
	if err != nil {
		log.Printf("Error decoding data from %s: %v\n", db.filename, err)
		return err
	}

	return nil

}

func (db *KeyValueDB) save() error {
	// open the database file
	file, err := os.Create(db.filename)
	if err != nil {
		log.Printf("Error creating file %s: %v\n", db.filename, err)
		return err
	}
	defer file.Close()

	// Lock the file for writing
	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX); err != nil {
		log.Printf("Error locking file %s for writing: %v\n", db.filename, err)
		return err
	}
	defer syscall.Flock(int(file.Fd()), syscall.LOCK_UN) // Unlock when done

	// encode the database file
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(db.data)
	if err != nil {
		log.Printf("Error encoding data to %s: %v\n", db.filename, err)
		return err
	}

	return nil
}

func (db *KeyValueDB) Set(key, value string) error {
	// lock the database
	db.Lock()
	defer db.Unlock()

	// add key-value pair to database
	dbEntry := db.data[key]
	dbEntry.Value = value
	dbEntry.Timestamp = append(dbEntry.Timestamp, time.Now())
	db.data[key] = dbEntry

	// save the database
	db.save()
	return nil
}

func (db *KeyValueDB) Get(key string) (string, error) {
	// lock the database for reading
	db.RLock()
	defer db.RUnlock()

	// get the value for the key
	dbEntry, found := db.data[key]
	if !found {
		return "", fmt.Errorf("key %s not found", key)
	}

	return dbEntry.Value, nil
}

func (db *KeyValueDB) Delete(key string) bool {
	db.Lock()
	defer db.Unlock()

	// delete the key-value pair from the database
	_, found := db.data[key]
	if found {
		delete(db.data, key)
		db.save()
	}

	return found
}

func (db *KeyValueDB) Timestamp(key string) ([]time.Time, error) {
	db.RLock()
	defer db.RUnlock()

	// get the value for the key
	DBEntry, found := db.data[key]
	if !found {
		return nil, fmt.Errorf("Key %s not found", key)
	}

	return DBEntry.Timestamp, nil

}

func (db *KeyValueDB) Close() error {
	return db.save()
}
