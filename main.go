/*
Requirements:
• solution should be written on Go;
• code should be commented in English;
• source data should be taken from stdin and results should go to stdout and be informative.
Task:
A man has a rather old car being worth $2000. He saw a secondhand car being worth $8000. He wants to keep
his old car until he can buy the secondhand one.
He thinks he can save $1000 each month but the prices of his old car and of the new one decrease of 1.5
percent per month. Furthermore the percent of loss increases by a fixed 0.5 percent at the end of every two
months.
Can you help him? Our man finds it difficult to make all these calculations.
How many months will it take him to save up enough money to buy the car he wants, and how much money will
he have left over?
*/

package main

import (
	"fmt"
)

func main() {
	var (
		savings                     float64
		oldCarPrice                 float64 = 2000
		newCarPrice                 float64 = 8000
		monthlySavings              float64 = 1000
		monthlyNegativeInterest     float64 = 0.015
		negativeInterestRaise       float64 = 0.005 //Every given period monthly loss percent will rise by that value
		negativeInterestRaisePeriod int     = 2     //Value in months
	)

	Calculations(savings, oldCarPrice, newCarPrice, monthlySavings,
		monthlyNegativeInterest, negativeInterestRaise, negativeInterestRaisePeriod)
}


/*
I'm aware it's not the best idea to store the money values as floats but unfortunately, there is no decimal type in Go
and I struggled to find an appropriate package to work with currency.
For instance, github.com/rhymond/go-money seemed to be a solution at the first glance, but it turned out it wouldn't
help with interest calculations anyway as its division and multiplication functions work solely with ints.
*/
func Calculations(savings, oldCarPrice, newCarPrice, monthlySavings,
monthlyNegativeInterest, negativeInterestRaise float64, negativeInterestRaisePeriod int) {
	month := 1

	for ; savings < newCarPrice - oldCarPrice; month++ {
		savings += monthlySavings
		newCarPrice *= 1- monthlyNegativeInterest
		oldCarPrice *= 1- monthlyNegativeInterest
		if month %negativeInterestRaisePeriod == 0{
			monthlyNegativeInterest += negativeInterestRaise
		}
		fmt.Printf("Month: %d\nSavings: $%5.2f New car price: $%5.2f Old car price: $%5.2f\n\n",
			month, savings, newCarPrice, oldCarPrice)
	}
	fmt.Printf("It will take a man %d months to save money for the new car. " +
		"It will cost him $%5.2f and he will be left with $%5.2f of his savings",
		month-1, newCarPrice, savings+oldCarPrice-newCarPrice)
}
