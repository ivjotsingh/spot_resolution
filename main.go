package main

import "fmt"

// defining a struct of car
type car struct {
	registerationNumber string
	color               string
}

//embedding car with slot
type slot struct {
	car        car
	slotNumber int
}

// embedding slots in parking lot
type parkingLot struct {
	slots             map[string]slot
	totalNumberOfCars int
}

func main() {

	car1 := car{"1234", "Red"}
	car2 := car{"2345", "Green"}

	slot1 := slot{car1, 1}
	slot2 := slot{car2, 2}
	parkingLot1 := new(parkingLot)
	parkingLot1.totalNumberOfCars = 2
	parkingLot1.slots = make(map[string]slot)
	parkingLot1.slots["1"] = slot1
	parkingLot1.slots["2"] = slot2

	for no, slot := range parkingLot1.slots {
		fmt.Println(no)
		fmt.Println(slot.car.color)
	}

}
