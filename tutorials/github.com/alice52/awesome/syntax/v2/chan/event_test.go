package _chan

import (
	"sync"
	"testing"
	"time"
)

func TestEvent(t *testing.T) {
	eventChannel := make(chan Event)
	var wg sync.WaitGroup

	// Start a goroutine to handle events
	wg.Add(1)
	go eventHandler(eventChannel, &wg)

	// Publish some events
	publishEvent(eventChannel, "Event 1")
	publishEvent(eventChannel, "Event 2")
	publishEvent(eventChannel, "Event 3")

	time.Sleep(100 * time.Second)

	// Close the channel to signal that no more events will be published
	close(eventChannel)

	// Wait for the event handler goroutine to finish
	wg.Wait()
}
