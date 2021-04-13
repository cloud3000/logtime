package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
)

type clockConf struct {
	SEC string `json:"sec"`
	MIN string `json:"min"`
	HR  string `json:"hr"`
}

func readConf(fname string) clockConf {
	// json data
	var obj clockConf

	// read file
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Print(err)
		return obj
	}

	// Unmarshal json.
	// NOTE: When we edit clock.json with VsCode we get a json error, but the Unmarshal is successful.
	// ALSO NOTE: Some editors do not write to the real file until you exit.
	// For best results use VI to edit the clock.json file.
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
		return obj
	}

	return obj
}

func watcher(wg *sync.WaitGroup, dir string, fname string, done <-chan bool, chg chan bool) {
	defer wg.Done()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					if fname == event.Name {
						chg <- true
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("watcher error:", err)
			}
		}
	}()
	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("watcher has started watching %s.\n", dir)
	<-done
	log.Println("watcher is shuting down.")
}
