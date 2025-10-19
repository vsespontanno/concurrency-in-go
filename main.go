package main

func main() {
	channel := make(chan int)
	channel <- 1
	println(<-channel)
}
