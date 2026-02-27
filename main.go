package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello! Go egg test started ✅")
	for {
		fmt.Println("Tick:", time.Now().Format(time.RFC3339))
		time.Sleep(5 * time.Second)
	}
}