package main 

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)


var transRe = regexp.MustCompile(`(\d+) shares of ([A-Z]+) for \$(\d+(\.\d+)?)`)


type Transaction struct {
	Symbol string
	Volume int
	Price float64
}

func parseLine(line string) (Transaction, error) {
	matches := transRe.FindStringSubmatch(line)
	if matches == nil {
			return Transaction{}, fmt.Errorf("bad line: %q", line)
	}
	var transac Transaction
	transac.Symbol = matches[2]
	transac.Volume, _ =strconv.Atoi(matches[1])
	transac.Price, _ =strconv.ParseFloat(matches[3],64)
	
	return transac, nil

}


func main() {
	line := "12 shares of MSFT for $234.57"
	t, err := parseLine(line)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", t)
}