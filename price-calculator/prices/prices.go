package prices

import (
	"fmt"

	"majoramari.com/price-calculator/conversion"
	"majoramari.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"text_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tex_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)

		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteJson(job)
}

func New(fm *filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   *fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
