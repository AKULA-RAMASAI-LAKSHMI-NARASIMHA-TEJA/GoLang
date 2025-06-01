package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now) // 2025-05-31 22:44:50.5986025 +0530 IST m=+0.000525601

	// "01-02-2006" is used to format date as standard format according to document
	// "01-02-2006 Monday" for date day
	// "01-02-2006 15:04:05 Monday" for date time day"
	fmt.Println(now.Format("01-02-2006")) // 05-31-2025

	// time.Date() is used to create a specific date

}
