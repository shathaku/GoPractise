package main

import (
	"fmt"
	"sync"
)

type Income struct{
	Source string
	Amount int
}

var wg sync.WaitGroup

func main() {
	var bankBalance int
	var balance sync.Mutex

	incomes := []Income{
		{Source: "Main Job", Amount: 1000},
		{Source: "Gifts", Amount: 200},
		{Source: "Part Time Job", Amount: 500},
		{Source: "Stocks", Amount: 100},
	}

	wg.Add(200)
	// add income for next 52 weeks
	for i:=0; i<200; i++ {

		go func(wg *sync.WaitGroup){
			defer wg.Done()

			for _, income := range incomes {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
			}
		}(&wg)
	}
	
	wg.Wait()

	fmt.Printf("Final bank balance: %d", bankBalance)
}