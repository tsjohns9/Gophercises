package db

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

// Task is a task
type Task struct {
	Key       int
	Value     string
	Completed bool
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
		num := itob(key)
		mItem := marshal(int(key), task, false)
		return b.Put(num, mItem)
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
			item := unmarshal(v)
			tasks = append(tasks, item)
		}

		return nil
	})
	return tasks, err
}

// Delete removes an item by key
func Delete(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		bts := itob(uint64(key))
		return b.Delete(bts)
	})
}

// Complete marks a task as completed
func Complete(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		btsInt := itob(uint64(key))
		valueInBytes := bucket.Get(btsInt)
		task := unmarshal(valueInBytes)
		mItem := marshal(key, task.Value, true)
		return bucket.Put(btsInt, mItem)
	})
}

func itob(num uint64) []byte {
	bts := make([]byte, 8)
	binary.BigEndian.PutUint64(bts, num)
	return bts
}

func unmarshal(value []byte) Task {
	var item Task
	json.Unmarshal(value, &item)
	return item
}

func marshal(key int, value string, completed bool) []byte {
	item := Task{
		Key:       key,
		Value:     value,
		Completed: completed,
	}
	js, _ := json.Marshal(item)
	return js
}
