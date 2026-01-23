package week2

func WordCount(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++

	}
	return counts
}

//  Write a function that counts how many times each word appears.

// Rules

// Use a map[string]int

// No external packages

// Case-sensitive

// Empty slice → empty map

// Function signature
// func WordCount(words []string) map[string]int

// Example
// input := []string{"go", "rust", "go", "go", "java"}
// output := map[string]int{
// 	"go":   3,
// 	"rust": 1,
// 	"java": 1,
// }

// Bonus constraints (optional but respected 😈)

// Do NOT check if the key exists explicitly

// Idiomatic Go only
