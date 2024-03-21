package main

import (
	"fmt"
	"time"
)
import "github.com/robfig/cron/v3"

type job struct {
}

func (job) Run() {
	fmt.Println(time.Now(), "I am job1")
}

// https://pkg.go.dev/github.com/robfig/cron#section-readme
// Funcs are invoked in their own goroutine, asynchronously.
func main() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0/2 * * * * ?", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m1s", func() { fmt.Println("Every hour thirty") })

	c.AddJob("@every 1s", &job{})

	c.Start()

	// Funcs may also be added to a running Cron
	c.AddFunc("0/2 * * * * ?", func() { fmt.Println("Every day") })

	// Inspect the cron job entries' next and previous run times.
	// fmt.Printf("%v", c.Entries())

	time.Sleep(3 * time.Second)

	c.Stop() // Stop the scheduler (stop all jobs).

	time.Sleep(10000000000000)
}
