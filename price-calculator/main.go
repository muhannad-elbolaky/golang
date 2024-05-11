package main

import (
	"fmt"

	"majoramari.com/price-calculator/filemanager"
	"majoramari.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.New(fm, taxRate)
		priceJob.Process()
	}
}
