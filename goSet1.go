package main

import "fmt"

func main() {
	//savings
	grossPay := 100000
	taxRate := 0.2
	expenses := 30000
	remaining := savings(grossPay, taxRate, expenses)
	fmt.Printf("Savings after tax and expenses: %d centavos\n", remaining)

	// materialWaste
	totalMaterial := 100
	materialUnits := "kg"
	numJobs := 3
	jobConsumption := 25
	waste := materialWaste(totalMaterial, materialUnits, numJobs, jobConsumption)
	fmt.Printf("Material remaining after jobs: %s\n", waste)

	// interest
	principal := 50000
	rate := 0.05
	periods := 2
	finalValue := interest(principal, rate, periods)
	fmt.Printf("Final investment value: %d centavos\n", finalValue)

	//shiftletter
	fmt.Println("Testing shiftLetter")
	fmt.Printf("shiftLetter(\"A\", 0) = %q\n", shiftLetter("A", 0)) // A
	fmt.Printf("shiftLetter(\"A\", 2) = %q\n", shiftLetter("A", 2)) // C
	fmt.Printf("shiftLetter(\"Z\", 1) = %q\n", shiftLetter("Z", 1)) //  A
	fmt.Printf("shiftLetter(\"X\", 5) = %q\n", shiftLetter("X", 5)) // C
	fmt.Printf("shiftLetter(\" \", 5) = %q\n", shiftLetter(" ", 5)) // BLANK

	// caesarCipher
	fmt.Println("Testing caesarCipher")
	fmt.Printf("caesarCipher(\"HELLO WORLD\", 3) = %q\n", caesarCipher("HELLO WORLD", 3)) // KHOOR ZRUOG

	// vigenereCipher https://www.wikiwand.com/en/articles/Vigen%C3%A8re_cipher
	fmt.Println("Testing vigenereCipher ")
	fmt.Printf("vigenereCipher(\"A C\", \"KEY\") = %q\n", vigenereCipher("A C", "KEY"))     // K A
	fmt.Printf("vigenereCipher(\"HELLO\", \"KEY\") = %q\n", vigenereCipher("HELLO", "KEY")) // RIJVS

	// scytaleCipher
	fmt.Println("Testing scytaleCipher")
	message := "INFORMATION_AGE"
	encoded := scytaleCipher(message, 3)
	fmt.Printf("Original: %q\n", message)
	fmt.Printf("Encoded: %q\n", encoded) // Should be "IMNNA_FTAOIGROE"

}

// Savings calculates the money remaining for an employee after taxes and expenses.
//
// To get the take-home pay of an employee:
// 1. Apply the tax rate to the gross pay of the employee, round down.
// 2. Subtract the expenses from the after-tax pay of the employee.
//
// Params:
// - grossPay, the gross pay of an employee for a certain time period, expressed in centavos
// - taxRate, the tax rate for a certain time period, expressed as a number between 0 and 1 (e.g., 0.12)
// - expenses, the expenses of an employee for a certain time period, expressed in centavos
//
// Returns:
// - the number of centavos remaining from an employee's pay after taxes and expenses
func savings(grossPay int, taxRate float64, expenses int) int {
	afterTax := int(float64(grossPay) * (1 - taxRate))
	return afterTax - expenses
}

// MaterialWaste calculates how much material input will be wasted after running a number of jobs that consume a set amount of material.
//
// To get the waste of a set of jobs:
// 1. Multiply the number of jobs by the material consumption per job
// 2. Subtract the total material consumed from the total material available
//
// Format the output as a string, annotated with the units in which the material is expressed. Do not add a space between the number and the unit.
//
// Params:
// - totalMaterial, how much material you have at the start
// - materialUnits, the unit used to express an amount of material, e.g., "kg"
// - numJobs, how many jobs to run
// - jobConsumption, how much material each job consumes
//
// Returns:
// - the amount of remaining material expressed with its unit (e.g., "10kg")
func materialWaste(totalMaterial int, materialUnits string, numJobs int, jobConsumption int) string {
	consumed := numJobs * jobConsumption
	remaining := totalMaterial - consumed
	return fmt.Sprintf("%d%s", remaining, materialUnits)
}

// Interest calculates the final value of an investment after gaining simple interest over a number of periods.
//
// To get sample interest, simply multiply the principal to the quantity (rate * time). Add this amount to the principal to get the final value.
//
// Round down the final amount.
//
// Params:
// - principal, the principal, or starting, amount invested, expressed in centavos
// - rate, the interest rate per period, expressed as a decimal representation of a percentage (e.g., 3% is 0.03)
// - periods, the number of periods invested
//
// Returns:
// - the final value of the investment
func interest(principal int, rate float64, periods int) int {
	interest := float64(principal) * rate * float64(periods)
	return principal + int(interest)
}
