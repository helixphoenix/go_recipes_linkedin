package main 

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

)





func go_check(r io.Reader)(map[string]int, error){


	matches := make(map[string]int)
	s:= bufio.NewScanner(r)
	for s.Scan(){
		if strings.Contains(s.Text(), ";go "){
			s := strings.Split(s.Text(), " ")
				matches[s[2]]++
		}
	}
	if err := s.Err(); err != nil {
			return nil, err
	}

	return matches, nil

}



func main() {

		zsh, err := os.Open("Chapter_3/zsh_history.txt")
		if err !=nil {
			log.Fatal(err)
		}
		defer zsh.Close()

		matches, err := go_check(zsh)
		if err != nil {
				log.Fatal(err)
		}
		for command, count := range matches{
			fmt.Println( "'go",command,"' command is happening", count, "times")

		}



}