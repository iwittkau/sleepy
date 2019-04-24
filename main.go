package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

var (
	usage = `  usage: sleepy duration
	 %s
	 example: sleepy 10s (sleeps for 10 seconds)
`
	validUnits = "valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(usage, validUnits)
		os.Exit(1)
	}

	d, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(validUnits)
		os.Exit(1)
	}

	if d.Nanoseconds() < time.Second.Nanoseconds() {
		time.Sleep(d)
		os.Exit(0)
	}

	count := d.Nanoseconds() / time.Second.Nanoseconds()
	bar := pb.StartNew(int(count))
	bar.ShowCounters = false
	bar.ShowTimeLeft = true
	bar.SetUnits(pb.U_DURATION)
	for i := 0; i < int(count); i++ {
		bar.Increment()
		time.Sleep(time.Second)
	}

}
