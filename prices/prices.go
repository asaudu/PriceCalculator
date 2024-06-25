package prices

import (
	"fmt"

	"addyCodes.com/price-calculator/conversion"
	"addyCodes.com/price-calculator/prices/iomanager"
)

type TaxIncludedPriceJob struct {
	//implementing struct tags, "-" means exclude
	IOManager iomanager.IOManager `json:"-"`
	//the other struct tags simply format how the label appears
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}
	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	// the below line would overwrite the 1 result file with each job, so intead
	// filemanager.WriteJSON("result.json", job)
	//we implement fmt.Sprintf so we can generate a file with each job instead
	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30, 40},
		TaxRate:     taxRate,
	}
}
