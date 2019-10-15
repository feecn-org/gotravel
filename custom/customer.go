package custom

import (
	"fmt"
)

//channel message
var messages chan string = make(chan string)
var mapChan = make(chan map[string]string)
var currentHashMap = make(map[string]string)

func MakeCustomer() {
	currentHashMap["currentState"] = "goPro" //os.Getenv("GOENV");
	//fmt.Println(currentHashMap["currentState"])
	//go mapChan <- currentHashMap
	go func(mapChannel map[string]string) {
		mapChan <- mapChannel
	}(currentHashMap)
	go func(message string) {
		messages <- message
	}("Ping!")
	goPro()
	makeMapChannel()
}

func goPro() {
	fmt.Println(<-messages)
}
func makeMapChannel() {
	mapChannel := <-mapChan
	fmt.Println(mapChannel["currentState"])
	fmt.Println(mapChannel)
}
func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}
