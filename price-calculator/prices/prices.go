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

// LoadData loads data from the IOManager and populates the InputPrices field of the TaxIncludedPriceJob struct.
//
// It reads the lines from the IOManager using the ReadLines method and converts the lines to floats using the StringsToFloat function from the conversion package.
// If any error occurs during the process, it prints the error message and returns.
// Otherwise, it assigns the converted prices to the InputPrices field of the TaxIncludedPriceJob struct.
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

// Process calculates the tax-included prices for the input prices based on the tax rate and saves the results.
//
// No parameters.
// No return values.
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

// New creates a new TaxIncludedPriceJob instance with the provided FileManager and tax rate.
//
// Parameters:
// - fm: a pointer to a FileManager instance.
// - taxRate: a float64 representing the tax rate.
// Returns:
// - *TaxIncludedPriceJob: a pointer to the newly created TaxIncludedPriceJob instance.
func New(fm *filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   *fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
