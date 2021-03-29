package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))

	// convert time zone
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	datetime := time.Now().In(loc)
	fmt.Println(datetime)
	fmt.Println(datetime.Year())
	fmt.Println(datetime.Month())
	fmt.Println(datetime.Day())
	fmt.Println(datetime.Hour())
	fmt.Println(datetime.Minute())
	fmt.Println(datetime.Second())

	// timestamp
	fmt.Println(datetime.Unix())
	fmt.Println(datetime.UnixNano())
}
