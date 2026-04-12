package history

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	bolt "go.etcd.io/bbolt"
)

const bucket = "history"

// Entry represents a single history record.
type Entry struct {
	ID     string    `json:"id"`
	Type   string    `json:"type"` // ask | explain | run
	Query  string    `json:"query"`
	Result string    `json:"result"`
	Time   time.Time `json:"time"`
}

// dbPath returns the path to the history database (~/.pilot/history.db).
func dbPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".pilot")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "history.db"), nil
}

// open opens (or creates) the bbolt database and ensures the bucket exists.
func open() (*bolt.DB, error) {
	path, err := dbPath()
	if err != nil {
		return nil, err
	}
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1})
	if err != nil {
		return nil, err
	}
	return db, db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
}

// Save stores a query and its result in the history database.
func Save(typ, query, result string) error {
	db, err := open()
	if err != nil {
		return err
	}
	defer db.Close()

	e := Entry{
		ID:     fmt.Sprintf("%d", time.Now().UnixNano()),
		Type:   typ,
		Query:  query,
		Result: result,
		Time:   time.Now(),
	}
	data, _ := json.Marshal(e)

	return db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Put([]byte(e.ID), data)
	})
}

// List returns the most recent n entries, newest first.
func List(limit int) ([]Entry, error) {
	db, err := open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var entries []Entry
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucket)).Cursor()
		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			var e Entry
			if json.Unmarshal(v, &e) == nil {
				entries = append(entries, e)
			}
			if len(entries) >= limit {
				break
			}
		}
		return nil
	})
	return entries, nil
}

// Search returns entries whose query contains the given keyword (case-insensitive).
func Search(keyword string) ([]Entry, error) {
	all, err := List(500)
	if err != nil {
		return nil, err
	}
	var results []Entry
	kw := strings.ToLower(keyword)
	for _, e := range all {
		if strings.Contains(strings.ToLower(e.Query), kw) {
			results = append(results, e)
		}
	}
	return results, nil
}

// Clear deletes all history entries.
func Clear() error {
	db, err := open()
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(bucket))
	})
}