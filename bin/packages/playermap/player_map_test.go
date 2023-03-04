package playermap

import (
	"testing"
)

func TestPlayerMapCreated(t *testing.T) {
  _, error := New()

  if error != nil {
    t.Fatal("Constructor FAILED! Got a bad PlayerMap object.")
  }
}

func TestPlayerMapSizings(t *testing.T) {
  playerMap, _ := New()

  height := len(playerMap.theMap)
  width := len(playerMap.theMap["A"])

  if height != width || height != 10 || width != 10 {
    t.Fatal("The sizes of the map are different and not 10.")
  }
}

func TestPlayerMapConstructorReturnsEmptyMap(t *testing.T) {
  playerMap, _ := New()

  for _, arrayOfLetters := range playerMap.theMap {
    for _, letter := range arrayOfLetters {
      if letter != "E" {
        t.Fatal("Brand new player map is *NOT* fully empty.")
      }
    }
  }
}

func TestGetCellOfTheMapMethod(t *testing.T) {
   playerMap, _ := New()

  playerMap.GetCell("A", 1)

  if playerMap.GetCell("A", 1) != "E" {
    t.Fatal("The GetCell() function returned non-empty cell.")
  }

}

func TestMapDamageMethod(t *testing.T) {
  playerMap, _ := New()

  playerMap.GetDamage("A", 1)

  if playerMap.GetCell("A", 1) != "X" {
    t.Fatal("The damage did not set by the method GetDamage.")
  }
}

func TestMapPlaceUnit(t *testing.T) {
  playerMap, _ := New()

  playerMap.PlaceUnit("B", 3)

  if playerMap.GetCell("B", 3) != "S" {
    t.Fatal("The unit's cells not placed.")
  }
}
