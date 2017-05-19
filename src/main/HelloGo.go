package main

import (
	"fmt"

	"cn/org/xiwi/designpattern/compute"

	"cn/org/xiwi/designpattern/factory"
)

func main() {
	var compute compute.Computer = compute.Computer{
		Num1: 10,
		Num2: 5,
	}

	compute.SetStrategy(factory.NetStragegy("s"))

	fmt.Println(compute.Do())
}
