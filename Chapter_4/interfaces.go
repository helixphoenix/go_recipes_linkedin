package main 


import (
		"fmt"
)


type Thermostat struct {
		id string
		value float64
}


func (t *Thermostat) ID() string {
	return t.id
}

func (t *Thermostat) Value() float64 {
		return t.value
}


func (*Thermostat) Kind() string {
	return "thermostat"
}

type Camera struct {
		id string
}

func (c *Camera) ID() string {
		return c.id
}

func (*Camera) Kind() string {
		return "camera"
}

type Sensor interface {
		ID() string
		Kind() string
}

func printAll(sensors []Sensor) {
	for _,s := range sensors {
			fmt.Printf("Location: %s, kind: %s\n",s.ID(), s.Kind())
	}
}

func main() {
	t:= Thermostat{"Living Room", 25.13}
	c := Camera{"Cat room"}

	sensors:= []Sensor{&t, &c}
	printAll(sensors)
}