package db

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

// Task is a task
type Task struct {
	Key   int
	Value string
}

var db *bolt.DB

func init() {
	var err error
	db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}

// Create an item in the db
func Create(task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		key, _ := b.NextSequence()
		num := intToBytes(key)
		return b.Put(num, []byte(task))
	})

}

// List returns all items
func List() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(taskBucket)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			data := binary.BigEndian.Uint64(k)
			tasks = append(tasks, Task{Key: int(data), Value: string(v)})
		}

		return nil
	})
	return tasks, err
}

// Delete removes an item by key
func Delete(key int) error {

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		bts := intToBytes(uint64(key))
		return b.Delete(bts)
	})
}

func intToBytes(num uint64) []byte {
	bts := make([]byte, 8)
	binary.BigEndian.PutUint64(bts, num)
	return bts
}
