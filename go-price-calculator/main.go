package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, rate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)

		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", rate*100))

		// cmdManager := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(fileManager, rate)
		go priceJob.Process(doneChans[i], errorChans[i])

		// if err != nil {
		// 	fmt.Println("Could not process job,", err)
		// }
	}

	for i := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println("Could not process job,", err)
			}
		case <-doneChans[i]:
			fmt.Println("Done")
		}
	}

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
}
