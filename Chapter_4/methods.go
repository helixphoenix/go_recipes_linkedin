package main

import "fmt"

type Thermostat struct {
	ID    string
	value float64
}

func (t *Thermostat) Value() float64 {
	return t.value
}

func (t *Thermostat) Set(value float64) {
	 t.value = value
}

func (*Thermostat) Kind() string {
		return "thermostat"
}

func main() {
		t := Thermostat{"Living Room", 25}
		fmt.Printf("%s before: %.2f\n", t.ID, t.Value())
		t.Set(35)
		fmt.Printf("%s after: %2.f\n", t.ID, t.Value())
}