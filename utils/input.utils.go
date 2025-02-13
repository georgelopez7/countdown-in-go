package utils

import (
	"bufio"
	"fmt"
	"os"
)

func CheckUserInput() bool {
	// --- PROMPT USER BEFORE WRITING TO FILE ---
	fmt.Print("Would you like to save the solutions to a file? (Y/n): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	if userInput == "y" || userInput == "Y" {
		return true
	}

	fmt.Println("Solutions will not be saved.")
	return false
}
