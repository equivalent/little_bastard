package main

import "os"
import "strconv"
import "fmt"
import "time"
import "net/http"

func main() {
	if os.Getenv("URL") == "" {
		os.Exit(3)
	} else {
		for {
			var sleeptime int
			if os.Getenv("SLEEPFOR") == "" {
				sleeptime = 1000
			} else {
				sleeptime, _ = strconv.Atoi(os.Getenv("SLEEPFOR"))
			}

			fmt.Println("Request to " + os.Getenv("URL"))
			fmt.Println(http.Get(os.Getenv("URL")))
			time.Sleep(time.Duration(sleeptime) * time.Millisecond)
		}
	}
}
