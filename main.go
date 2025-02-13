package main

import (
	"fmt"
	"gofigure/utils"
	"strings"
	"sync"
)

func main() {
	// --- CREATE COUNTDOWN GAME ---
	numbers, target := utils.GenerateNumbers()
	fmt.Printf("\nNumbers: %v\n", numbers)
	fmt.Printf("Target: %d\n\n", target)

	// --- STORE PERMUTATIONS & SOLUTIONS ---
	permutations := utils.GetPermutations(numbers)
	solutionsChn := make(chan string, 1000)
	var wg sync.WaitGroup

	// --- PARALLEL SOLVING ---
	for _, p := range permutations {
		wg.Add(1)
		go func(p []int) {
			defer wg.Done()
			var steps string // STEPS TO REACH SOLUTION

			found := utils.SolveViaRecursion(p, target, &steps)
			if found {
				solutionsChn <- fmt.Sprintf("Solution found:\n%s", steps)
			}

		}(p)
	}

	// --- WAIT FOR ALL GOROUTINES TO FINISH ---
	go func() {
		wg.Wait()
		close(solutionsChn)
	}()

	// --- PREPARE TO HANDLE SOLUTIONS ---
	var solutionsBuffer strings.Builder
	uniqueSolutions := make(map[string]bool)
	uniqueSolutionCount := 0

	fmt.Printf("Finding solutions...\n")

	// --- HANDLE SOLUTIONS AS THEY COME DOWN THE CHANNEL ---
	for solution := range solutionsChn {
		if _, exists := uniqueSolutions[solution]; !exists {
			uniqueSolutions[solution] = true
			uniqueSolutionCount++
			solutionsBuffer.WriteString(solution + "\n")
		}
	}

	// --- PRINT NUMBER OF UNIQUE SOLUTIONS ---
	fmt.Printf("Number of unique solutions: %d\n\n", uniqueSolutionCount)

	// --- CHECK USER INPUT ---
	shouldContinue := utils.CheckUserInput()
	if !shouldContinue {
		return
	}

	// --- CREATE SOLUTIONS TEXT FILE ---
	utils.CreateTextFile("solutions.txt", solutionsBuffer.String())
	fmt.Println("Solutions has been saved!")

}
