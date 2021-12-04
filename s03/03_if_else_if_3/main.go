package main

import "fmt"

func main() {
	serverStatusOk := true

	if serverStatusOk {
		fmt.Println("Server is up & running!")
	}

	if serverName := "XF"; serverName == "XF" && serverStatusOk {
		fmt.Printf("%s server is up & running!", serverName)
	}

	//fmt.Println(serverName) // compiler error: undeclered variable
}
