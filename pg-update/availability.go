// Inserts carpark availability API into postgres
package main

import (
	"encoding/json"
	"fmt"
	"github.com/iody/carpark/repo"
	"net/http"
	"strconv"
)

type Info struct {
	TotalLots     string `json:"total_lots"`
	LotType       string `json:"lot_type"`
	LotsAvailable string `json:"lots_available"`
}
type Data struct {
	Info           []Info `json:"carpark_info"`
	CarparkNumber  string `json:"carpark_number"`
	UpdateDatetime string
}
type Items struct {
	Data []Data `json:"carpark_data"`
}
type CarPark struct {
	Items []Items `json:"items"`
}

func callHttpAvailability() CarPark {
	request, err := http.NewRequest("GET", "https://api.data.gov.sg/v1/transport/carpark-availability", nil)

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	var cp CarPark
	json.NewDecoder(resp.Body).Decode(&cp)
	fmt.Printf("%+v\n", cp)
	return cp
}

func main() {
	cp := callHttpAvailability()
	for _, d := range cp.Items[0].Data {

		lt, err := strconv.Atoi(d.Info[0].TotalLots)
		if err != nil {
			panic(err)
		}
		la, err := strconv.Atoi(d.Info[0].LotsAvailable)
		if err != nil {
			panic(err)
		}

		a := repo.HdbCarparkAvailability{
			CarParkNo:     d.CarparkNumber,
			LotType:       d.Info[0].LotType,
			TotalLots:     int32(lt),
			LotsAvailable: int32(la),
		}

		repo.Insert(a)
	}
}
