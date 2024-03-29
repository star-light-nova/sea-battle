package ship

import (
    "fmt"
    "errors"
)

type Ship struct {
    class ShipClass;
    CellsToOccupy ShipOccupy;
}

func New(class ShipClass) (*Ship, error) {
    _, isShipExists := shipsAndOccupyCells[class]

    if !isShipExists {
        return nil, errors.New("The ship of this class does not exist.")
    }

    return &Ship {
        class: class,
        CellsToOccupy: shipsAndOccupyCells[class],
    }, nil
}

func (ship Ship) String() string {
    return fmt.Sprintf("The ship: %s, occupies: %d", ship.class, ship.CellsToOccupy)
}
