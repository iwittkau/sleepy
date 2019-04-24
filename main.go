package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

var (
	usage = `  usage: sleepy duration [message]

	 %s
	
	 examples: 
	 	  sleepy 10s 
	          	sleeps for 10 seconds
		   
		  sleepy 10s "sleeping 10 seconds" 
		  	sleeps for 10 seconds and displays 
			"sleeping 10 seconds" as the progress bar prefix
			and a message when sleeping has finished
`
	validUnits = "valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(usage, validUnits)
		os.Exit(1)
	}

	var prefix string
	if len(os.Args) > 2 {
		prefix = fmt.Sprintf(" %s ", os.Args[2])
		log.SetFlags(log.Ltime | log.Ldate)
	}

	defer printJobName(prefix)

	d, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(validUnits)
		os.Exit(1)
	}

	if d.Nanoseconds() < time.Second.Nanoseconds() {
		time.Sleep(d)
		return
	}

	count := d.Nanoseconds() / time.Second.Nanoseconds()
	bar := pb.StartNew(int(count))
	bar.Prefix(prefix)
	bar.ShowCounters = false
	bar.ShowTimeLeft = true
	bar.SetUnits(pb.U_DURATION)
	for i := 0; i < int(count); i++ {
		bar.Increment()
		time.Sleep(time.Second)
	}
}

func printJobName(s string) {
	if s != "" {
		s = strings.TrimSpace(s)
		log.Printf("finished %q\n", s)
	}
}
