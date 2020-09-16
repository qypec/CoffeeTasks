// Complete the solution so that it reverses the string passed into it.

package main

func Solution(word string) string {
	result := make([]rune, len(word))

	i := len(word) - 1
	for _, letter := range word {
		result[i] = rune(letter)
		i--
	}
	return (string(result))
}
