package main

// Logging Time, could be the heart-beat of a scheduling system.

import (
	"log"
	"sync"
	"time"
)

//  watchFile contains the logMessage.SEC, logMessage.MIN, and logMessage.HR.
var watchPath = "config"
var watchFile = "config/clock.json"

func main() {
	// readConf, channels, and watcher are all in support of realtime message changes.
	logMessage := readConf(watchFile)
	newConf := make(chan bool)
	quit := make(chan bool)

	wg := new(sync.WaitGroup) // I want to wait for watcher to end.
	wg.Add(1)                 // I know I don't need to wait, I'm doing it just because I should.

	go watcher(wg, watchPath, watchFile, quit, newConf)

	// The ticker and timing variables are for logging time.
	clockTick := time.NewTicker(time.Second)
	min := time.Minute
	hr := time.Hour
	deadline := hr * 3
	msg := logMessage.SEC

	// This process will run until the duration of sec >= deadline.

	for sec := time.Second; sec <= deadline; sec = sec + time.Second {

		select {
		case <-clockTick.C:
			// Set the value of msg, depending on the modulus of time duration.
			if sec%min == 0 {
				msg = logMessage.MIN
			}
			if sec%hr == 0 {
				msg = logMessage.HR
			}
			// Print msg, and set default msg
			log.Printf("%s, total time: %v \n", msg, sec)
			msg = logMessage.SEC

		case <-newConf: // Fires whenever clock.json is changed
			logMessage = readConf(watchFile)
		}

		// If passed deadline, tell the watcher to quit, log the event, and stop the ticker.
		if sec >= deadline {
			quit <- true
			log.Printf("Deadline of %v has passed.\n", sec)
			clockTick.Stop()
			break
		}

	}
	// Wait for the watcher to stop, it's the right thing to do ;-)
	wg.Wait()
	println("Main is shutting down.")
}
