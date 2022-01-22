package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.UTC)
	fmt.Println(time.Local)

	now := time.Now()
	fmt.Println(now)
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())

	timestamp := now.Unix()
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	fmt.Println(timeObj.Year())

	fmt.Println(now.Weekday().String())
}
