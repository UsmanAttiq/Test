package main

import (
    "fmt"
    "os"
)

func main() {
	x:=0
	var P,S,F int =1526, 9583, 37575
	for x<1{
		fmt.Print("Please select an option:\n")
		fmt.Print("1 – Print Covid19 cases in Pakistan\n") 
		fmt.Print("2 – Print Covid19 cases in South Korea\n") 
		fmt.Print("3 – Print Covid19 cases in France\n") 
		fmt.Print("4 – Print my personalized message to Coronavirus\n") 
		fmt.Print("0 - Exit\n")
	    fmt.Print("Option: ")
	    var input string
	    _, err := fmt.Scanln(&input)
	    if err != nil {
	        fmt.Fprintln(os.Stderr, err)
	        return
	    }
	    switch input {
	    case "1":
	        fmt.Println("Total Cases: ",P)
	        fmt.Print("\n\n\n")
	    case "2":
	        fmt.Println("Total Cases: ",S)
	        fmt.Print("\n\n\n")
	    case "3":
	        fmt.Println("Total Cases: ",F)
	        fmt.Print("\n\n\n")
	    case "4":
	        fmt.Println("Just when I was starting to actively strive for what I wanted from life. Corona Strikes. Or maybe this is an excuse, kidding. Stay safe in Quarantine!")
	        fmt.Print("\n\n\n")
	    case "0":
	        x++
	    }
	}
}