package main

import (
	"fmt"
	"io"
	"log"
	"os"
)


type Trade struct {
	Symbol string
	Volume int
	Price float64

}


func generateReport(w io.Writer, trades []Trade){
	for i, t := range trades {
			log.Printf("%d %#v",i,t)
			fmt.Fprintf(w, "%-10s %04d %.2f\n", t.Symbol,t.Volume, t.Price)
	}
}

func main() {
	log.SetPrefix("LOG: ")

	trades := []Trade{
		{"MSFT", 231, 234.57},
		{"TSLA", 123, 686.75},
		{"BRK-A",1, 399100},
	}
	generateReport(os.Stdout, trades)
}