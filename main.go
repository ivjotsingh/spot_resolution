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
	car        map[string]string
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
		parking.slots[i] = Slot{make(map[string]string), i}
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
		if len(parkingLot.slots[k].car) == 0 {
			emptySlot = k
			break
		}
	}
	return emptySlot
}

/*
    def free_slot(self, slot_number):
        if not self.is_parking_lot_functional():
            return parking_lot_not_functional

        try:
            slot = self.slots[str(slot_number)]
            if slot.car:
                self.registration_numbers_of_cars_parked.remove(slot.car.registration_number)
                slot.car = None
                return f'Slot number {slot_number} is free'
            else:
                return f'Slot number {slot_number} was already free'
        except KeyError:
			return f'Slot number {slot_number} does not exist in the parking lot'
*/

func (parkingLot ParkingLot) freeSlot(slotNumber int) {
	fmt.Println(parkingLot.slots[slotNumber])
	parkingLot.slots[slotNumber] = Slot{make(map[string]string), slotNumber}
	fmt.Println(parkingLot.slots[slotNumber])
}

func (parkingLot ParkingLot) parkCar(car Car) {
	nearestParkingSlot := parkingLot.getNearestParkingSlot()
	parkingLot.slots[nearestParkingSlot].car[car.registerationNumber] = car.color

}

func main() {
	totalNumberOfCars := 5
	parkingLot := createParkingLot(totalNumberOfCars)
	fmt.Println(parkingLot)
	parkingLot.parkCar(Car{registerationNumber: "1234", color: "Red"})
	parkingLot.parkCar(Car{registerationNumber: "2345", color: "Green"})
	fmt.Println(parkingLot)
	parkingLot.freeSlot(1)
	fmt.Println(parkingLot)
	parkingLot.parkCar(Car{registerationNumber: "5345", color: "Green"})
	fmt.Println(parkingLot)
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
