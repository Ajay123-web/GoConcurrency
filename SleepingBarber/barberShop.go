package main

import "time"

type BarberShop struct {
	Open         bool
	ShopCapacity int
	HaircutTime  time.Duration
	TotalBarbers int
	BarbersDoneChan chan bool
	ClientsChan chan string
}