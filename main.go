package main

import (
	"fmt"

	"port-scanner.com/port"
)

func main() {

	fmt.Println("Port Scanner in Go")

	open := port.ScanPorts("tcp", "localhost", 1313)
	fmt.Printf("Port open: %t\n", open)

	results := port.InitialScan("localhost")
	fmt.Println(results)
}
