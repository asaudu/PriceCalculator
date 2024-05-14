package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30, 40}
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	var result map[float64][]float64 = make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxIncludedPries := make([]float64, len(prices))

		for priceIndex, price := range prices {
			taxIncludedPries[priceIndex] = price * (1 + taxRate)
		}

		result[taxRate] = taxIncludedPries
	}

	fmt.Println("Here's the result", result)
}
