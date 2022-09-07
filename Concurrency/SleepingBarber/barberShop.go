package main

import (
	"fmt"
	"time"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NoOfBarbers int
	BarbersDoneChan chan bool
	CustomersChan chan string
	ShopOpen bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NoOfBarbers++
	fmt.Printf("%s has entered the shop\n", barber)
	go func() {
		isSleeping := false

		for {
			if len(shop.CustomersChan) == 0 {
				fmt.Printf("%s goes to sleep\n", barber)
				isSleeping = true
			}

			customer, isOpen := <-shop.CustomersChan

			if isOpen {
				if isSleeping {
					isSleeping = false
					fmt.Printf("%s wakes up %s\n", customer, barber)
				}
				shop.cutHair(barber, customer)
			} else {
				shop.sendBarberHome(barber)
				return
			}
		}
		
	}()
}

func (shop * BarberShop) cutHair(barber, customer string) {
	time.Sleep(2*time.Second)
	fmt.Printf("%s cuts hair for %s\n", barber, customer)
}

func (shop * BarberShop) sendBarberHome(barber string) {
	fmt.Printf("%s goes back home\n", barber)
	shop.BarbersDoneChan <-true
}

func (shop *BarberShop) closeShop() {
	close(shop.CustomersChan)
	shop.ShopOpen = false


	for a:=1; a<=shop.NoOfBarbers; a++ {
		<- shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)
	fmt.Printf("Barber shop is closed for the day")
}

func (shop * BarberShop) addCustomer(customer string) {
	fmt.Printf("%s enters the shop\n", customer)

	if(shop.ShopOpen) {
		select {
		case shop.CustomersChan <- customer:
			fmt.Printf("%s takes a seat in waiting room\n", customer)
		default:
			fmt.Println("Waiting room is full !!")
		}
	} else {
		fmt.Println("shop is laready closed!!")
	}
	
}