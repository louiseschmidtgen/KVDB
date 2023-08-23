package database

import (
	"encoding/gob"
	"os"
	"sync"
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
	file, err := os.Open(db.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // KVDB does not exist yet
		}
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	return decoder.Decode(&db.data)
}

func (db *KeyValueDB) save() error {
	file, err := os.Create(db.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(db.data)
}

func (db *KeyValueDB) Set(key, value string) {
	db.Lock()
	defer db.Unlock()

	info := db.data[key]
	info.Value = value
	info.Timestamp = append(info.Timestamp, time.Now())
	db.data[key] = info

	db.save()
}

func (db *KeyValueDB) Get(key string) (string, []time.Time) {
	db.RLock()
	defer db.RUnlock()

	info, found := db.data[key]
	if !found {
		return "", nil
	}
	return info.Value, info.Timestamp
}

func (db *KeyValueDB) Delete(key string) bool {
	db.Lock()
	defer db.Unlock()

	_, found := db.data[key]
	if found {
		delete(db.data, key)
		db.save()
	}
	return found
}

func (db *KeyValueDB) Timestamp(key string) []time.Time {
	db.RLock()
	defer db.RUnlock()

	info, found := db.data[key]
	if !found {
		return nil
	}
	return info.Timestamp
}

func (db *KeyValueDB) Close() error {
	return db.save()
}
