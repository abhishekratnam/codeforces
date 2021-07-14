package main

import (
	"fmt"
	"strings"
)

// Write a program which prompts the user to enter a string. The program searches through the entered string for the characters â€˜iâ€™, â€˜aâ€™, and â€˜nâ€™. The program should print â€œFound!â€ if the entered string starts with the character â€˜iâ€™, ends with the character â€˜nâ€™, and contains the character â€˜aâ€™. The program should print â€œNot Found!â€ otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

// Examples: The program should print â€œFound!â€ for the following example entered strings, â€œianâ€, â€œIanâ€, â€œiuiygaygnâ€, â€œI d skd a efju Nâ€. The program should print â€œNot Found!â€ for the following strings, â€œihhhhhnâ€, â€œinaâ€, â€œxianâ€.

func main() {
	var stringInput string
	fmt.Println(`Please enter a string value with double quotes, ex "ian"`)
	fmt.Scanf("%q", &stringInput)

	hasPrefixI := strings.HasPrefix(stringInput, "i")
	hasSuffixN := strings.HasSuffix(stringInput, "n")
	hasContainA := strings.Contains(stringInput, "A")
	hasContaina := strings.Contains(stringInput, "a")

	if hasPrefixI && hasSuffixN && (hasContainA || hasContaina) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
