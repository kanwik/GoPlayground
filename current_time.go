package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Printf("Текущее время: %s\n", currentTime.Format("15:04:05"))
}
