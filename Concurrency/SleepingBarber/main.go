package main

import (
	"fmt"
	"time"
)

var seatingCapacity = 10

func main() {

	customersChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create barber shop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: (1000 * time.Millisecond),
		NoOfBarbers:     0,
		BarbersDoneChan: doneChan,
		CustomersChan:   customersChan,
		ShopOpen:        true,
	}

	// add a barber
	shop.addBarber("Frank")
	shop.addBarber("Bob")

	// add clients

	shopClosing := make(chan bool)
	closed := make(chan bool)

	// close shop after a while
	go func() {
		<-time.After(10*time.Second)
		shopClosing <- true
		shop.closeShop()
		closed <- true
	}()

	// add customer

	i := 1
	go func() {
		for {
			select {
			case <- shopClosing:
				return
			case <-time.After(time.Millisecond * 500):
				shop.addCustomer(fmt.Sprintf("Customer #%d", i))
				i++
			}
		}
	}()

	<- closed
	// time.Sleep(5 * time.Second)
}
