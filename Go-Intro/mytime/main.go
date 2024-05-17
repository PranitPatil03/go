package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Time")

	currentTime := time.Now()
	fmt.Println(currentTime)

	CurrentDate := currentTime.Format("01-02-2006")
	CurrentTime := currentTime.Format("15:-04:05")
	CurrentDay := currentTime.Format("Monday")

	fmt.Println(CurrentDate)
	fmt.Println(CurrentTime)
	fmt.Println(CurrentDay)

}
