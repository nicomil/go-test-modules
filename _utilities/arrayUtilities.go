package utilities

import "fmt"

// capital case function's names are exported, are like public
func PrintArr(arr []string) {
	for _, elem := range arr {
		print(elem)
	}
}

// lowercase function's names are not exported, are like private
func print(elem string) {
	fmt.Println(elem)
}
