package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var taskBucket = []byte("tasks")

// Task is a task
type Task struct {
	Key   int
	Value string
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		// task := strings.Join(args, " ")
		db, err := bolt.Open("./my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		updateErr := db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists(taskBucket)
			b := tx.Bucket(taskBucket)
			key, _ := b.NextSequence()
			fmt.Println(key)
			// b.Put(key, []byte(task))
			return err
		})
		if updateErr != nil {
			log.Fatal(updateErr)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
