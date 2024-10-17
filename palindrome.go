package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	// Remove spaces and convert to lowercase
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)

	// Check if the string is a palindrome
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}
