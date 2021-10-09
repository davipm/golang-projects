package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p("Time Now:", now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("Then:", then)

	p("Year:", then.Year())
	p("Month:", then.Month())
}
