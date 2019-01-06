package main

import (
	"flag"
	"time"
)

func main() {
	year := time.Now().Year()
	flag.IntVar(&year, "y", year, "Year, four digits")
	month := int(time.Now().Month())
	flag.IntVar(&month, "m", month, "Month, 1-12")
	flag.Parse()
}
