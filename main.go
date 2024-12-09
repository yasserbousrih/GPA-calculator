package main

import (
	"fmt"
)

func main() {
	var (
		numSubjects  int
		totalPoints  float64
		totalCredits float64
	)

	fmt.Print("Enter the number of subjects: ")
	fmt.Scanln(&numSubjects)

	for i := 1; i <= numSubjects; i++ {
		var gradePoints, credits float64
		fmt.Printf("Enter grade points for subject %d: ", i)
		fmt.Scanln(&gradePoints)
		fmt.Printf("Enter credits for subject %d: ", i)
		fmt.Scanln(&credits)

		totalPoints += gradePoints * credits
		totalCredits += credits
	}

	if totalCredits == 0 {
		fmt.Println("Total credits cannot be zero.")
		return
	}

	gpa := totalPoints / totalCredits
	fmt.Printf("Your GPA is: %.2f\n", gpa)
}
