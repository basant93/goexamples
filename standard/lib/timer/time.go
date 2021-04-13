package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("elapsed time: ", elapsed)

	today := time.Now()
	fmt.Println(today)

	todayFormated := today.Format("Apr 4, 2021 12:30:21")
	fmt.Println(todayFormated)
	futureDate := time.Date(2021, 12, 2, 0, 0, 0, 0, time.UTC)
	remainingTime := time.Until(futureDate)
	fmt.Println("remaining time ", remainingTime)
	fmt.Println(remainingTime.Hours())
	fmt.Println(remainingTime.Minutes())

	newDate := futureDate.AddDate(10, 1, 1)
	fmt.Println(newDate)
}
