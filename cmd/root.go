package cmd

import (
	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"
	"fmt"
	"encoding/json"
	"time"
	"log"
	"runtime"
	"path"
)


var db *bolt.DB
var open bool

func Open() error {
	var err error
	_, filename, _, _ := runtime.Caller(0)  // get full path of this file
	dbfile := path.Join(path.Dir(filename), "data.db")
	config := &bolt.Options{Timeout: 1 * time.Second}
	db, err = bolt.Open(dbfile, 0600, config)
	if err != nil {
		log.Fatal(err)
	}
	open = true
	return nil
}

func Close() {
	open = false
	db.Close()
}

type Task struct {
	Name string
	Complete bool
}

func (t *Task) add() error {
	err := db.Update(func(tx *bolt.Tx) error {
		tasks, err := tx.CreateBucketIfNotExists([]byte("task"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		enc, err := t.encode()
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		err = tasks.Put([]byte(t.Name), enc)
		return err
	})
	return err
}

func listTasks() {
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("task")).Cursor()
		i := 1
		for k, v := c.First(); k != nil; k, v = c.Next() {
			t, err := decodeTask(v)
			if err != nil {
				fmt.Println("error!")
			}
			if !t.Complete {
				fmt.Printf("%d - %s\n", i, t.Name)
				i++
			}
		}
		return nil
	})
}

func (t *Task) encode() ([]byte, error){
	enc, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func decodeTask(data []byte) (*Task, error) {
	var task *Task
	err := json.Unmarshal(data, &task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

var RootCmd = &cobra.Command{
	Use:   "task [command]",
	Short: "A task list",
}
