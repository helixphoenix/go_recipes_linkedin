package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Record struct {
	Time    time.Time
	Station string
	Temp    float64 `json:"temperature"`
	Rain    float64
}

func readRecord(r io.Reader) (Record, error) {
	var rec Record
	dec := json.NewDecoder(r)
	fmt.Println(dec)
	if err := dec.Decode(&rec); err != nil {
		return Record{}, err
	}
	return rec, nil
}

func main() {
	file, err := os.Open("Chapter_5/record.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rec, err := readRecord(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", rec)
}
