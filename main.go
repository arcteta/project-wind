package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

func main() {

	var windstatus string
	var waterstatus string

	for {
		valueWind := rand.Intn(101)
		valueWater := rand.Intn(101)

		if valueWater < 5 {
			waterstatus = "aman"
		} else if valueWater >= 5 && valueWater <= 8 {
			waterstatus = "siaga"
		} else if valueWater >= 9 {
			waterstatus = "Bahaya"
		} else {
			waterstatus = "invalid input"
		}

		if valueWind < 6 {
			windstatus = "aman"
		} else if valueWind >= 6 && valueWind <= 15 {
			windstatus = "siaga"
		} else if valueWind >= 6 {
			windstatus = "Bahaya"
		} else {
			windstatus = "invalid input"
		}

		Post(valueWater, valueWind, windstatus, waterstatus)
		time.Sleep(15 * time.Second)
	}

}

func Post(water int, wind int, swnd string, swtr string) {

	var data Data
	data.Water = water
	data.Wind = wind

	reqjson, err := json.Marshal(data)

	client := http.Client{}

	if err != nil {
		log.Fatal(err)
		return
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqjson))
	req.Header.Set("Conotent-Type", "application/json")

	if err != nil {
		log.Fatal(err)
		return
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatal(err)
	}

	tojson, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(tojson))
	fmt.Println("status wind:", swnd)
	fmt.Println("status water:", swtr)

}
