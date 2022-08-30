package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arrivalRate = 100
var duration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Cyan("The Sleeping Barber Problem")
	color.Cyan("____________________________")
	fmt.Println("")

	clientChan := make(chan string , seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop {
		Open: true,
		ShopCapacity: seatingCapacity,
		HaircutTime: duration,
		TotalBarbers: 0,
		BarbersDoneChan: doneChan,
		ClientsChan: clientChan,
	}

	color.Green("Shop open for day!!!")
	fmt.Println("")

	shop.AddBarber("Frank")

	time.Sleep(5 * time.Second)

}