package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iody/carpark/repo"
	"log"
	"net/http"
	"strconv"
)

type Carpark struct {
	Address       string  `json:"address"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
	TotalLots     int32   `json:"total_lots"`
	AvailableLots int32   `json:"available_lots"`
}

func carparkHandler(w http.ResponseWriter, r *http.Request) {
	latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 32)

	if err != nil {
		w.WriteHeader(400)
		return
	}
	longitude, err := strconv.ParseFloat(r.FormValue("longitude"), 32)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	page, err := strconv.Atoi(r.FormValue("page"))
	perPage, err := strconv.Atoi(r.FormValue("per_page"))

	hca, err := repo.GetNearest(float32(latitude), float32(longitude), perPage, page)
	if err != nil {
		panic(err)
	}

	var carparks = make([]Carpark, len(hca))
	for i, a := range hca {
		carparks[i] = Carpark{
			a.HdbCarparkInformation.Address,
			a.HdbCarparkInformation.Latitude,
			a.HdbCarparkInformation.Longitude,
			a.TotalLots,
			a.LotsAvailable,
		}
	}
	fmt.Println("Endpoint Hit: carparks")
	json.NewEncoder(w).Encode(carparks)
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/carparks/nearest", carparkHandler)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
