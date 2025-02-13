package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var largeNumbers = []int{25, 50, 75, 100}
var smallNumbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func GenerateNumbers() ([]int, int) {

	// --- CREATE RANDOM NUMBER GENERATOR ---
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// --- CHOOSE 0-4 LARGE NUMBERS ---
	numLarge := rng.Intn(5)
	chosenNumbers := []int{}

	// --- SHUFFLE LARGE NUMBERS AND PICK UNIQUE ONES ---
	shuffledLarge := append([]int(nil), largeNumbers...)
	rand.Shuffle(len(shuffledLarge), func(i, j int) { shuffledLarge[i], shuffledLarge[j] = shuffledLarge[j], shuffledLarge[i] })

	// --- PICK LARGE NUMBERS ---
	chosenNumbers = append(chosenNumbers, shuffledLarge[:numLarge]...)

	// --- SELECT REMAINING NUMBERS FROM SMALL NUMBERS ---
	for len(chosenNumbers) < 6 {
		index := rng.Intn(len(smallNumbers))
		chosenNumbers = append(chosenNumbers, smallNumbers[index]) // APPEND
	}

	// --- GENERATE A TARGET BETWEEN 100 AND 999 ---
	target := rng.Intn(900) + 100

	return chosenNumbers, target
}

func GetPermutations(arr []int) [][]int {
	var res [][]int
	var helper func([]int, int)

	helper = func(a []int, n int) {
		if n == 1 {
			res = append(res, append([]int(nil), a...))
			return
		}
		for i := 0; i < n; i++ {
			helper(a, n-1)
			if n%2 == 1 {
				a[0], a[n-1] = a[n-1], a[0]
			} else {
				a[i], a[n-1] = a[n-1], a[i]
			}
		}
	}

	helper(arr, len(arr))
	return res
}

func SolveViaRecursion(numbers []int, target int, steps *string) bool {

	// --- STEP 1: CHECK IF ANY NUMBER ALREADY MATCHES THE TARGET ---
	for _, num := range numbers {
		if num == target {
			return true
		}

	}

	// --- STEP 2: TRY ALL PAIRS OF NUMBERS TO GENERATE NEW NUMBERS ---
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {

			newNumbers := append([]int{}, numbers...)
			newNumbers = append(newNumbers[:j], newNumbers[j+1:]...) // Remove j
			newNumbers = append(newNumbers[:i], newNumbers[i+1:]...) // Remove i

			a, b := numbers[i], numbers[j]

			// --- STEP 3: PERFORM ALL VALID OPERATIONS ON THE PAIR ---
			operations := []struct {
				result int
				op     string
			}{
				{a + b, fmt.Sprintf("%d + %d = %d", a, b, a+b)}, // Addition (always valid)
				{a * b, fmt.Sprintf("%d * %d = %d", a, b, a*b)}, // Multiplication (always valid)
			}

			// OPERATION: a - b (Only valid if a >= b)
			if a >= b {
				operations = append(operations, struct {
					result int
					op     string
				}{a - b, fmt.Sprintf("%d - %d = %d", a, b, a-b)})
			}

			// OPERATION: b - a (Only valid if b >= a)
			if b >= a {
				operations = append(operations, struct {
					result int
					op     string
				}{b - a, fmt.Sprintf("%d - %d = %d", b, a, b-a)})
			}

			// OPERATION: a / b (Only valid if b != 0 AND a % b == 0)
			if b != 0 && a%b == 0 {
				operations = append(operations, struct {
					result int
					op     string
				}{a / b, fmt.Sprintf("%d / %d = %d", a, b, a/b)})
			}

			// OPERATION: b / a (Only valid if a != 0 AND b % a == 0)
			if a != 0 && b%a == 0 {
				operations = append(operations, struct {
					result int
					op     string
				}{b / a, fmt.Sprintf("%d / %d = %d", b, a, b/a)})
			}

			// --- STEP 4: RECURSIVELY TRY EACH NEW RESULT ---
			for _, op := range operations {
				newNumbers = append(newNumbers, op.result)

				if SolveViaRecursion(newNumbers, target, steps) {
					*steps = op.op + "\n" + *steps
					return true
				}

				// --- UNDO THE LAST OPERATION ---
				newNumbers = newNumbers[:len(newNumbers)-1]
			}
		}
	}

	// --- NO SOLUTION FOUND ---
	return false
}
