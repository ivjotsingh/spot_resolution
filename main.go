package main

import (
	"fmt"
	"sort"
)

// Car defining a struct of car
type Car struct {
	registerationNumber string
	color               string
}

// Slot embedding car with slot
type Slot struct {
	car        Car
	slotNumber int
}

// ParkingLot embedding slots in parking lot
type ParkingLot struct {
	slots             map[int]Slot
	totalNumberOfCars int
}

func createParkingLot(totalNumberOfCars int) ParkingLot {

	parking := new(ParkingLot)
	parking.slots = make(map[int]Slot)

	parking.totalNumberOfCars = totalNumberOfCars
	for i := 1; i <= totalNumberOfCars; i++ {
		parking.slots[i] = Slot{Car{}, i}
	}
	fmt.Println("created parking lot with slots: ", totalNumberOfCars)
	return *parking

}

func (parkingLot ParkingLot) getNearestParkingSlot() int {
	var keys []int
	var emptySlot int
	for k := range parkingLot.slots {
		keys = append(keys, k)

	}
	sort.Ints(keys)
	for _, k := range keys {
		if parkingLot.slots[k].car == (Car{}) {
			emptySlot = k
			break
		}
	}
	return emptySlot
}

func (parkingLot ParkingLot) parkCar(car Car) {
	nearestParkingSlot := parkingLot.getNearestParkingSlot()
	fmt.Println("printing")
	fmt.Println(parkingLot.slots[nearestParkingSlot].car)
	parkingLot.slots[nearestParkingSlot].car = Car

}

// func (p ParkingLot) getNearestParkingSlot() int (){
// 	var keys []int
// 	for k := range p.slots {
//         keys = append(keys, k)
// 	}
// 	sort.Ints(keys)

// 	return 1
// }

func main() {
	totalNumberOfCars := 5
	parkingLot := createParkingLot(totalNumberOfCars)
	fmt.Println(parkingLot)
	parkingLot.parkCar(Car{})
}

/*
func main() {

	totalNumberOfCars := 5
	parking := createParkingLot(totalNumberOfCars)
	// for no, slot := range parking.slots {
	// 	if slot.car == (Car{}) {
	// 		fmt.Println(slot.slotNumber)
	// 		fmt.Println(no)
	// 	}
	// }
	parking.getNearestParkingSlot()

	/*
		for _, slot := range parking.slots {
			fmt.Println(slot.slotNumber)
			fmt.Println(slot)
		}
*/

/*
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
*/

//}
