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

import "fmt"

func main() {
	const kilogramsPerPound = 0.45359237
	const metersPerInch = 0.0254

	var weight float64
	fmt.Print("Enter weight in pounds: ")
	fmt.Scanf("%f ", &weight)

	var height float64
	fmt.Print("Enter height in inch: ")
	fmt.Scanf("%f ", &height)

	// Convert Pounds to Kilograms and Inches to Meters
	weightInKilograms := weight * kilogramsPerPound
	heightInMeters := height * metersPerInch

	bmi := weightInKilograms / (heightInMeters * heightInMeters)

	weightStatus := ""
	switch true {
	case bmi < 18.5:
		weightStatus = "Underweight"
	case bmi < 25:
		weightStatus = "Normal"
	case bmi < 30:
		weightStatus = "Overweight"
	default:
		weightStatus = "Obese"
	}

	fmt.Println()
	fmt.Printf("%-10s %-10s %-12s %-10s \n", "Weight", "Height", "BMI", "Status")
	fmt.Printf("%-10.2f %-10.2f %-12.6f %-10s \n", weight, height, bmi, weightStatus)
}
