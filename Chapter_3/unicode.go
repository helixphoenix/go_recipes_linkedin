package main 

import (
	"fmt"
	"unicode/utf8"
)


func lineLength(words []string) int {
	total := 0

    for _, word := range words{
		total += utf8.RuneCountInString(word)
	}

	numSpace := len(words) -1
	return total + numSpace

}


func main() {
	words := []string{"<<", "Don;t", "Panic",">>"}
    fmt.Println(lineLength(words))
}