package repo

import (
	"fmt"
	"testing"
)

func TestNearest(t *testing.T) {
	res, err := GetNearest(1.37326, 103.897, 1, 10)

	if err != nil {
		panic(err)
	}
	if len(res) == 0 {
		t.Errorf("GetNearest ; want size > %d", len(res))
	}
	for _, a := range res {
		fmt.Println(a.HdbCarparkInformation.CarParkNo)
		if a.LotsAvailable < 0 {
			t.Errorf("GetNearest ; want lots_availbale < %d", 0)
		}
	}
}
