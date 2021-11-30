package main 

import (
	"fmt"
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)
func parseTime(ts string)(time.Time,error){
		return time.Parse("2006-01-02 15:04:05 PM",ts)
}

func laggingStations(r io.Reader, timeout time.Duration) ([]string, error){

    var reply struct {
		LastCheckTime string
		Stations []struct {
				Name string
				Status string 
				LastCheck struct {
						Time string 
				}
		}
	}
	 
	decod := json.NewDecoder(r)
	if err := decod.Decode(&reply); err != nil {
			return nil, err
	}
	fmt.Println(reply)

	checkTime, err := parseTime(reply.LastCheckTime)
	if err != nil {
		return nil, err
	}
	var lagging []string 
	for _, station := range reply.Stations {
			if station.Status != "Active" {
				continue
			}
			lastCheck, err := parseTime(station.LastCheck.Time)
			if err != nil {
				return nil, err
			}

			if checkTime.Sub(lastCheck) > timeout {
					lagging= append(lagging, station.Name)
			}
	}
	return lagging, nil
}


func main() {
	file, err := os.Open("Chapter_5/stations.json")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	delayed, err := laggingStations(file, time.Minute)
	if err != nil {
		log.Fatal(err)
	}
    
	for _, name := range delayed {
			fmt.Println(name)
	}


}