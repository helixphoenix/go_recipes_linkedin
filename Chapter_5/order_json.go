package main

import (
		"encoding/json"
		"fmt"
		"os"
)


type Quantity struct {
	Value float64
	Unit string 
}


func (q *Quantity) MarshalJSON() ([]byte , error) {
	if q.Unit == "" {
			return nil, fmt.Errorf("empty unit")
	}
	text := fmt.Sprintf("%f%s", q.Value, q.Unit)
	return json.Marshal(text)
 }


 func main(){
	 l := Quantity{1.89,"m"}
	 p := Quantity{35, "cm"}
     json.NewEncoder(os.Stdout).Encode(&l)
     json.NewEncoder(os.Stdout).Encode(&p)

 }