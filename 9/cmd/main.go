package main

import "fmt"

func main() {
	pit()
}

func pit() {
	for i := 0; i < 3; i++ {
		func() {
			fmt.Printf("%d ", i)
		}()
	}
	// Output: 3 3 3
}
