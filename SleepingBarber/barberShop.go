package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	Open         bool
	ShopCapacity int
	HaircutTime  time.Duration
	TotalBarbers int
	BarbersDoneChan chan bool
	ClientsChan chan string
}

func (shop *BarberShop) AddBarber(barber string) {
	shop.TotalBarbers++

	go func() {
		isSleeping := false

		color.Yellow("%s is checking the waiting room" , barber)

		for {
			if len(shop.ClientsChan) == 0 {
				isSleeping = true
				color.Red("%s is going to sleep" , barber)
			}
			//is shop is close we will close the channels and shopOpen will become false
			client , shopOpen := <-shop.ClientsChan
			if(shopOpen) {
				if isSleeping {
					color.Yellow("%s wakes %s to cut his hair" , client , barber)
					isSleeping = false
				}

				shop.CutHair(barber , client)

			} else {
				shop.SendBarberHome(barber)
				return
			}
		}

	}()
}

func (shop *BarberShop) CutHair(barber , client string) {
	color.Green("%s is cutting %s's hair" , barber , client)
	time.Sleep(shop.HaircutTime)
	color.Green("%s has finished cutting %s's hair" , barber , client)
}

func (shop *BarberShop) SendBarberHome(barber string) {
	color.Cyan("%s is going home" , barber)
	shop.BarbersDoneChan <- true
}