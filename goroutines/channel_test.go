package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Indra Mahessa"
	}()

	name := <-channel
	fmt.Println(name)
}

func ChannelAsParameter(channel chan string) {
		time.Sleep(2 * time.Second)
		channel <- "Indra Mahessa"
		fmt.Println("Hello World")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	
	go ChannelAsParameter(channel)

	name := <-channel
	fmt.Println(name)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Indra Mahesa"
}

func OnlyOut(channel <-chan string) {
	name := <-channel
	fmt.Println(name)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 1)
	defer close(channel)

	channel <- "Indra"
	
	fmt.Println(<-channel)
}


func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Indra Mahessa" + strconv.Itoa(i)
		}
		close(channel)
	}()
	
	for name := range channel {
		fmt.Println(name)
	}

	fmt.Println("Done")
}

func GiveResponseToChannel(channel chan<- string, data string, isDelayed bool) {
	// if(isDelayed) {
	// 	time.Sleep(2 * time.Second)
	// }

	channel <- data
}
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveResponseToChannel(channel1, "Indra", false)
	go GiveResponseToChannel(channel2, "Mahesa", true)

	counter := 0
	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dapat dari channel 1 " + data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dapat dari channel 2 " + data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter >= 2 {
			break
		}
	}

	fmt.Println("Selesai")
}