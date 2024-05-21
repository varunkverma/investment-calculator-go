package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
	// var years float64 = 10
	fmt.Printf("enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Printf("enter the years for investment: ")
	fmt.Scan(&years)

	fmt.Printf("enter the expected return rate for investment: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
