package main

import (
	"fmt"
	"github.com/UsmanAttiq/Test/covid_pkg"
)

func main() {
	fmt.Printf(covid.PrintPakistanCases())
	fmt.Printf("\n")
	fmt.Printf(covid.PrintSouthKoreaCases())
	fmt.Printf("\n")
	fmt.Printf(covid.PrintFranceCases())
	fmt.Printf("\n")
	fmt.Printf(covid.PrintMessage())
}
