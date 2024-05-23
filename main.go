package main

import (
	"fmt"

	"port-scanner.com/port"
)

func main() {

	fmt.Println("Port Scanner in Go")
	//fmt.Println(len(port.GetCommonPorts()))

	//open := port.ScanPorts("tcp", "localhost", 1313)
	//fmt.Printf("Port open: %t\n", open)

	// Normal scan
	//results := port.InitialScan("localhost")
	//fmt.Println(results)

	// Extensive Scan
	//extensiveResults := port.ExtensiveScan("localhost")
	//fmt.Println(extensiveResults)

	//quick Scan
	quickResults := port.QuickScan("localhost")
	fmt.Println(quickResults)
}
