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
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("A man has a rather old car being worth some money. He saw a secondhand car.\nHe wants to keep " +
		"his old car until he can buy the secondhand one.\n" +
		"He thinks he can save a particular sum of money each month but the prices of " +
		"his old car and of the new one decrease every month.\n" +
		"Furthermore the percent of loss increases by a fixed percent every several months.\n" +
		"To help a man do all the calculations, let's input specific values.\n")

	var (
		savings                     float64
		oldCarPrice                 float64 = 2000
		newCarPrice                 float64 = 8000
		monthlySavings              float64 = 1000
		monthlyNegativeInterest     float64 = 0.015
		negativeInterestRaise       float64 = 0.005 //Every given period monthly loss percent will rise by that value
		negativeInterestRaisePeriod int     = 2     //Value in months
	)

	fmt.Println("1. Input all the values manually.\n2. Calculate with default values.")

	if getPositiveFloat() == 1 {
		fmt.Println("\nHow much can a man save per month?")
		monthlySavings = getPositiveFloat()
		fmt.Println("\nWhat is the initial old car price?")
		oldCarPrice = getPositiveFloat()
		fmt.Println("\nWhat is the initial new car price?")
		newCarPrice = getPositiveFloat()
		fmt.Println("\nHow much value do cars lose every month? In other words, what is the negative interest? " +
			"Value should be between 0 and 1.")
		negativeInterestRaise = getInterest()
		fmt.Println("\nHow often a negative inerest rate inreses? Value in whole months.")
		negativeInterestRaisePeriod = getPositiveInt()
		fmt.Println("\nBy how much a negative inerest rate inreses? Value should be between 0 and 1.")
		negativeInterestRaise = getInterest()
	}

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

	for ; savings < newCarPrice-oldCarPrice; month++ {
		savings += monthlySavings
		newCarPrice *= 1 - monthlyNegativeInterest
		oldCarPrice *= 1 - monthlyNegativeInterest
		if month%negativeInterestRaisePeriod == 0 {
			monthlyNegativeInterest += negativeInterestRaise
		}
		fmt.Printf("Month: %d\nSavings: $%5.2f New car price: $%5.2f Old car price: $%5.2f\n\n",
			month, savings, newCarPrice, oldCarPrice)
	}
	fmt.Printf("It will take a man %d months to save money for the new car. "+
		"It will cost him $%5.2f and he will be left with $%5.2f of his savings",
		month-1, newCarPrice, savings+oldCarPrice-newCarPrice)
}

//returns a value between 0 and 1
func getInterest() (res float64) {
	res = getPositiveFloat()
	if res <= 0 || res >= 1 {
		fmt.Println("Please, enter a number between 0 and 1, both included")
		return getInterest()
	}
	return
}

func getPositiveFloat() (res float64) {
	_, err := fmt.Scanf("%f", &res)
	if err != nil{
		flushStdin()
		fmt.Println("Please, enter a number")
		return getPositiveFloat()
	} else if res < 0 {
		fmt.Println("Please, enter a positive number")
		return getPositiveFloat()
	} else{
		return
	}
}

func getPositiveInt() (res int) {
	_, err := fmt.Scanf("%d", &res)
	if err != nil{
		flushStdin()
		fmt.Println("Please, enter a whole number")
		return getPositiveInt()
	} else if res < 0 {
		fmt.Println("Please, enter a positive whole number")
		return getPositiveInt()
	} else{
		return
	}
}
func flushStdin() {
	stdin := bufio.NewReader(os.Stdin)
	stdin.ReadString('\n')
}
