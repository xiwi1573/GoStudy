package main

import "fmt"

func main() {
	fmt.Println(fab(34))
}

func fab(i int) int {
	if i < 2 {
		return i
	}
	return fab(i-2) + fab(i-1)
}
