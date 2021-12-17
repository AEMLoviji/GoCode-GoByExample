package main

func main() {
	c := make(chan string)
	c <- "No one likes my channel!" // fatal error: all goroutines are asleep - deadlock!
}
