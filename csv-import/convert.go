package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Converts the geo 3414 to 4326 format
func convert(x string, y string) map[string]float32 {
	request, err := http.NewRequest("GET",
		fmt.Sprintf("https://developers.onemap.sg/commonapi/convert/3414to4326?X=%s&Y=%s", x, y), nil)

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	d := map[string]float32{}
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", pos)
	return d
}

func loadCsv() {
	file, err := os.Open("data/hdb-carpark-information.csv")
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	out, err := os.Create("data/sample.csv")
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		fmt.Println("Error:", err)
		return
	}

	// automatically call Close() at the end of current method
	defer out.Close()

	// on the fly
	r := csv.NewReader(file)
	w := csv.NewWriter(out)

	r.Read() // skip 1st line
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var pos = convert(record[2], record[3])
		record[2] = fmt.Sprintf("%f", pos["latitude"])
		record[3] = fmt.Sprintf("%f", pos["longitude"])

		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}

		fmt.Println(record)
	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
func main() {
	loadCsv()
}
