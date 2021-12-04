/*
Exercise: Write a program  to calculate body mass index(BMI)

Read 'weight' (in pounds) and 'height' (in inches) from console

bmi = weight / (height * height )

bmi < 18.5 --> Underweight
bmi < 25   --> Normal
bmi < 30   --> Overweight
otherwise  --> Obese
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const minRandomVal = 10
	const maxRandomVal = 22

	var yourRandomVal int
	for {
		fmt.Printf("Enter a random number between %d and %d: ",
			minRandomVal, maxRandomVal)

		fmt.Scanf("%d ", &yourRandomVal)

		if (yourRandomVal < minRandomVal) || (yourRandomVal > maxRandomVal) {
			fmt.Println("*Incorrect input value(s). Please try again ...")
			continue
		}

		break
	}

	rand.Seed(time.Now().UnixNano())
	systemRandomValue := minRandomVal + rand.Intn(maxRandomVal-minRandomVal)

	yourRandomValSumOfDigits :=
		(yourRandomVal / 10) + (yourRandomVal % 10)
	systemRandomValSumOfDigits :=
		(systemRandomValue / 10) + (systemRandomValue % 10)

	lotteryResult := ""
	if systemRandomValue == yourRandomVal {
		lotteryResult = "You won $1,000"
	} else if yourRandomValSumOfDigits == systemRandomValSumOfDigits {
		lotteryResult = "You won $500"
	} else {
		lotteryResult = "Try Again!"
	}

	fmt.Printf("\nyourRandomVal=%d systemRandomValue=%d \nlotteryResult=%s \n",
		yourRandomVal, systemRandomValue, lotteryResult)

}
