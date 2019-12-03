// Repository to connect to pg
package repo

import (
	"fmt"
	"github.com/go-pg/pg"
	"os"
)

// Entity
type HdbCarparkAvailability struct {
	CarParkNo             string `pg:",pk"`
	TotalLots             int32
	LotType               string
	LotsAvailable         int32
	HdbCarparkInformation HdbCarparkInformation
	tableName             struct{} `pg:"hdb_carpark_availability"`
}

// Entity
type HdbCarparkInformation struct {
	CarParkNo                string `pg:",pk"`
	Address                  string
	Latitude                 float32  `pg:"x_coord"`
	Longitude                float32  `pg:"y_coord"`
	tableName                struct{} `pg:"hdb_carpark_information,discard_unknown_columns"`
	HdbCarparkAvailabilities []*HdbCarparkAvailability
}

var db *pg.DB

func Insert(ca HdbCarparkAvailability) {
	err := db.Insert(&ca)

	fmt.Println(ca.CarParkNo)
	fmt.Println(ca.LotsAvailable)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully inserted!")
}

// Find the nearest car park given a given point, offset and limit can be pass as well
func GetNearest(x float32, y float32, l int, o int) ([]HdbCarparkAvailability, error) {
	var cpa []HdbCarparkAvailability
	err := db.Model(&cpa).
		Relation("HdbCarparkInformation").
		Where("lots_available > ?", 0).
		OrderExpr("geom <-> ST_SetSRID(ST_MakePoint(?, ?), 4326)", y, x).
		Limit(l).
		Offset(o).
		Select()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully queried!")

	return cpa, nil
}

func init() {
	url := os.Getenv("DATABASE_URL")
	if len(url) == 0{
		url = "postgres://docker:docker@192.168.1.6/gis?sslmode=disable"
	}
	opt, err := pg.ParseURL(url)
	if err != nil {
		panic(err)
	}
	db = pg.Connect(opt)
}
