package main

import (
	"fmt"
	"sync"

	"BinanceParse/pkg/binance"
)

func main() {
	binanceClient := binance.NewClient()
	symbols, err := binanceClient.GetSymbols()
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	priceChan := make(chan map[string]float64)

	for _, symbol := range symbols {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			price, err := binanceClient.GetPrice(s)
			if err != nil {
				fmt.Println(err)
				return
			}
			priceChan <- map[string]float64{s: price}
		}(symbol)
	}

	go func() {
		wg.Wait()
		close(priceChan)
	}()

	for priceMap := range priceChan {
		for symbol, price := range priceMap {
			fmt.Printf("%s %.5f\n", symbol, price)
		}
	}
}
