package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v\n:", err)
		os.Exit(1)
	}
	fmt.Println("Current time: ", ntpTime.Format(time.TimeOnly))
}
