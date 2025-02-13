package utils

import (
	"fmt"
	"os"
)

func CreateTextFile(fileName string, fileContent string) {
	// --- CREATE FILE ---
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// --- WRITE CONTENT TO FILE ---
	_, err = file.WriteString(fileContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
