package mapandship

import (
	"testing"
	"github.com/StarLightNova/sea-battle/bin/packages/coordinates"
)

func TestMapShipUnitsPlacement(t *testing.T) {
  masi := New()
  coor, _ := coordinates.New("a1:b1")

  masi.placeShip(*coor)

  if masi.playerMap.GetCell("A", 1) != "S" && masi.playerMap.GetCell("B", 1) != "S" {
    t.Fatal("The ship did not placed on the passed coordinates.")
  }
}
