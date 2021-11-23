package main 

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func grep(r io.Reader, term string)([]string, error){
	var matches [] string
	s:= bufio.NewScanner(r)
	for s.Scan(){
		if strings.Contains(s.Text(), term){
			    matches = append(matches, s.Text())
		}
	}
	if err := s.Err(); err != nil {
			return nil, err
	}

	return matches, nil

}



func main() {

		journal, err := os.Open("Chapter_3/journal.txt")
		if err !=nil {
			log.Fatal(err)
		}
		defer journal.Close()

		matches, err := grep(journal, "System is rebooting")
		if err != nil {
				log.Fatal(err)
		}
		fmt.Printf("%d reboots\n", len(matches))
}