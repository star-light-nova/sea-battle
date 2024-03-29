package playersinit

import (
    "fmt"
    "testing"
)

func TestPlayerConstructor(t *testing.T) {
    _, error := New()

    if error != nil {
        t.Fatal("The players wrapper (initializer) didn't create an instance.")
    }
}

func TestPlayersMapsandShipPlacements(t *testing.T) {
    players, error := New()

    players.PutShips()

    fmt.Println("The first Player:\n", players.FirstPlayer)
    fmt.Println()
    fmt.Println("Second Player:\n", players.SecondPlayer)

    if error != nil {
        t.Fatal("The players map are incorrect.")
    }
}

func TestPlayersCurrentAndOpponentsMap(t *testing.T) {
    players, error := New()

    players.PutShips()
    
    fmt.Println(players.EntireMap())

    if error != nil {
        t.Fatal("Can't print the entire map.")
    }
}
