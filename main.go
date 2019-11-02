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
	slot              slot
	totalNumberOfCars int
}

func main() {

	car1 := car{"1234", "Red"}

	slot1 := slot{car1, 1}
	parkingLot1 := parkingLot{slot1, 1}

	fmt.Println(parkingLot1.slot.car.color)
}
