package main

import (
	"addyCodes.com/price-calculator/cmdmanager"
	"addyCodes.com/price-calculator/prices"
)

func main() {
	//prices := []float64{10, 20, 30, 40}
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}

}
