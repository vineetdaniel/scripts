package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"/../../backup"

	"github.com/matryer/filedb"
)

type path struct {
	Path string
	Hash string
}

func main() {
	m := &backup.Monitor{
		Destination: *archive,
		Archiver:    backup.ZIP,
		//we create a map ready to store the path and hashes internally
		Paths: make(map[string]string),
	}
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()
	var (
		interval = flag.Int("interval", 10, "interval between checks (seconds)")
		archive  = flag.String("archive", "archive", "path to archive location")
		dbpath   = flag.String("db", "./db", "path to filedb database")
	)
	flag.Parse()

	db, err := filedb.Dial(*dbpath)
	if err != nil {
		fatalErr = err
		return
	}
	defer db.Close()
	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}

	var path path
	col.ForEach(func(_ int, data []byte) bool {
		if err := json.Unmarshal(data, &path); err != nil {
			fatalErr = err
			return true
		}
		m.Paths[path.Path] = path.Hash
		return false
	})
	if fatalErr != nil {
		return
	}
	if len(m.Paths) < 1 {
		fatalErr = errors.New("no paths")
		return
	}
	check(m, col)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-time.After(time.Duration(*interval) * time.Second):
			check(m, col)
		case <-signalChan:
			//stop
			fmt.Println()
			log.Printf("stopping...")
			goto stop
		}
	}
stop:
}

func check(m *backup.Monitor, col *filedb.C) {
	log.Println("Checking....")
	counter, err := m.Now()
	if err != nil {
		log.Fatalln("failed to backup", err)
	}
	if counter > 0 {
		log.Printf(" Archived %d directories\n", counter)

		var path path
		col.SelectEach(func(_ int, data []byte) (bool, []byte, bool) {
			if err := json.Unmarshal(data, &path); err != nil {
				log.Println("failed to unmarshal data (skipping)", err)
				return true, data, false
			}
			path.Hash, _ = m.Paths[path.Path]
			newdata, err := json.Marshal(&path)
			if err != nil {
				log.Println("failded to marshal data (skipping)", err)
				return true, data, false
			}
			return true, newdata, false
		})
	} else {
		log.Println(" No changes")

	}
}
