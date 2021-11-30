package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"encoding/json"
	"math"
)

func parseTime(ts string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04", ts)
}

func findSpeed(t1, t2 string, distance float64) (float64, error) {
	start, err := parseTime(t1)
	if err != nil {
		log.Fatal(err)
	}
	end, err := parseTime(t2)
	if err != nil {
		log.Fatal(err)
	}

	dur := float64(end.Sub(start)) / float64(time.Hour)
	return float64(distance / dur), err

}



func maxRideSpeed(r io.Reader) (float64, error) {
	var ride_car struct {
		Rides []struct{
			StartTime string `json:"start"`
			EndTime   string `json:"end"`
			Distance  float64 `json:"distance"`
		}
	}
	decod := json.NewDecoder(r)
	if err := decod.Decode(&ride_car); err != nil {
			return 0, err
	}
	maxSpeed := float64(0)
    for _, ride := range ride_car.Rides{
			speed, err := findSpeed(ride.StartTime, ride.EndTime, ride.Distance)

			if err != nil {
				return 0, err
			}
			maxSpeed = math.Max(speed, maxSpeed)

	 }

	return maxSpeed, nil

}

func main() {
	file, err := os.Open("Chapter_5/rides.json")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	speed, err := maxRideSpeed(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(speed)
}
