package main

import "encoding/json"
import "os"
import "strconv"
import "fmt"
import "time"
import "net/http"

type UrlSleep struct {
	Sleep int    `json:"sleep"`
	Url   string `json:"url"`
}

type UrlsSleepsJSON struct {
	Urls []UrlSleep `json:"urls"`
}

func main() {
	loc, _ := time.LoadLocation("UTC")

	if os.Getenv("URLS") != "" {
		res := UrlsSleepsJSON{}
		json.Unmarshal([]byte(os.Getenv("URLS")), &res)

		lastExecutions := []time.Time{}
		for range res.Urls {
			lastExecutions = append(lastExecutions, time.Now().In(loc))
		}

		for {
			for y, x := range res.Urls {
				lastExecutions[y] = sleepyRequest(x.Url, loc, lastExecutions[y], x.Sleep)
			}
		}
	} else {
		if os.Getenv("URL") == "" {
			os.Exit(3)
		} else {
			lastExecution := time.Now().In(loc)

			var sleepFor int
			if os.Getenv("SLEEPFOR") == "" {
				sleepFor = 1000
			} else {
				sleepFor, _ = strconv.Atoi(os.Getenv("SLEEPFOR"))
			}

			res := UrlSleep{Url: os.Getenv("URL"), Sleep: sleepFor}

			for {
				lastExecution = sleepyRequest(res.Url, loc, lastExecution, res.Sleep)
			}
		}
	}
}

func sleepyRequest(url string, loc *time.Location, lastExecution time.Time, sleepFor int) time.Time {
	if time.Now().After(lastExecution.In(loc).Add(time.Duration(sleepFor) * time.Millisecond)) {
		fmt.Println("Request to " + url)
		fmt.Println(http.Get(url))
		return time.Now().In(loc)
	} else {
		return lastExecution.In(loc)
	}
}
