package main

import (
	"fmt"
	"strings"
)

type Letter struct {
	Symbol  string
	English string
}

var letters = []Letter{
	{"Σ", "Sigma"}, {"σ", "Fitifiti"}, {"ς", "Nice tries"},
}

func englishFor(greek string) (string, error) {
	for _, letter := range letters {
		if strings.Contains(greek, letter.Symbol) {
			return letter.English, nil
		}

	}
	return "", fmt.Errorf("unknown greek letter: %#v", greek)
}

func main() {
	fmt.Println(englishFor("Σ"))
	fmt.Println(englishFor("σ"))
	fmt.Println(englishFor("ς"))
}
