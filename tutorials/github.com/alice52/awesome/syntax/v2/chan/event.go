package _chan

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Message string
}

func publishEvent(ch chan<- Event, message string) {
	// Create and send an event to the channel
	event := Event{Message: message}
	ch <- event
}

func eventHandler(ch <-chan Event, wg *sync.WaitGroup) {
	defer wg.Done()

	// Loop to handle events until the channel is closed
	for {
		event, ok := <-ch
		if !ok {
			// Channel is closed, no more events to process
			fmt.Println("Event handler exiting.")
			return
		}

		// Process the event
		processEvent(event)
	}
}

func processEvent(event Event) {
	// Simulate event processing
	fmt.Printf("Processing event: %s\n", event.Message)
	time.Sleep(1 * time.Second) // Simulating some work
}
