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

	const lowRange = 1
	const highRange = 10
	const numberOfChanges = 5

	rand.Seed(time.Now().UnixNano())
	randomNo := lowRange + rand.Intn(highRange-lowRange)

	fmt.Printf("You have %d chances to guess "+
		"a system-generated random number.\n", numberOfChanges)

	var yourNum int
	message := "You were not lucky this time. Try again!"
	for i := 0; i < numberOfChanges; i++ {
		fmt.Printf("Enter a number between %d and %d try(#%d): ",
			lowRange, highRange, i+1)
		fmt.Scanf("%d ", &yourNum)

		if yourNum == randomNo {
			message = "You guessed the number!"
			break
		}

		if yourNum > randomNo {
			fmt.Printf("Pick a smaller number.")
		} else {
			fmt.Printf("Pick a bigger number.")
		}
	}

	fmt.Println("\n----------------------------------------------")
	fmt.Printf("-- Your Number=%d Random Number=%d \n-- %s",
		yourNum, randomNo, message)
	fmt.Println("\n----------------------------------------------")
}
