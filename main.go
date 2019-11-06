package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Car defining a struct of car
type Car struct {
	registerationNumber string
	color               string
}

// Slot embedding car with slot
type Slot struct {
	car        map[int]Car
	slotNumber int
}

// ParkingLot embedding slots in parking lot
type ParkingLot struct {
	slots              map[int]Slot
	totalNumberOfSlots int
}

// createParkingLot creates a parking lot object and assign slots to it
func createParkingLot(totalNumberOfSlots int) ParkingLot {

	parking := new(ParkingLot)
	parking.slots = make(map[int]Slot)

	parking.totalNumberOfSlots = totalNumberOfSlots
	for i := 1; i <= totalNumberOfSlots; i++ {
		parking.slots[i] = Slot{make(map[int]Car), i}
	}
	fmt.Println("created parking lot with slots: ", totalNumberOfSlots)
	return *parking

}

// getNearestParkingSlot gets nearest empty slot
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

func (parkingLot ParkingLot) parkCar(car Car) bool {
	nearestParkingSlot := parkingLot.getNearestParkingSlot()
	if nearestParkingSlot == 0 {
		fmt.Println("Parking is full")
		return false
	}
	parkingLot.slots[nearestParkingSlot].car[nearestParkingSlot] = car
	fmt.Println("Car", car.registerationNumber, "parked at:", nearestParkingSlot)
	return true
}

func (parkingLot ParkingLot) freeSlot(slotNumber int) bool {
	if len(parkingLot.slots[slotNumber].car) == 0 {
		fmt.Println("Slot Number:", slotNumber, "is already free")
		return false
	}
	parkingLot.slots[slotNumber] = Slot{make(map[int]Car), slotNumber}
	fmt.Println("Slot Number:", slotNumber, "is free")
	return true
}

func (parkingLot ParkingLot) searchCarByColor(color string) []int {
	parkingSlotsWithCars := []int{}
	for k, v := range parkingLot.slots {
		if v.car[v.slotNumber].color == color {
			parkingSlotsWithCars = append(parkingSlotsWithCars, k)
		}
	}
	fmt.Println("slot numbers of cars with color", color, parkingSlotsWithCars)
	return parkingSlotsWithCars

}

func (parkingLot ParkingLot) searchCarByRegisterationNumber(registerationNumber string) []int {
	parkingSlotsWithCars := []int{}
	for k, v := range parkingLot.slots {
		if v.car[v.slotNumber].registerationNumber == registerationNumber {
			parkingSlotsWithCars = append(parkingSlotsWithCars, k)

		}
	}
	fmt.Println("slot numbers of cars with Registeration Number", registerationNumber, parkingSlotsWithCars)
	return parkingSlotsWithCars
}

func (parkingLot ParkingLot) getParkingLotStatus() {
	for slotNumber, slot := range parkingLot.slots {
		if len(slot.car) != 0 {
			fmt.Println(slot.car[slotNumber].registerationNumber, "is parked at slot number:", slotNumber)
		}
	}
}

func main() {

	fmt.Println("Enter the no of slots in parking lot")
	var totalNumberOfSlots int
	fmt.Scanf("%d", &totalNumberOfSlots)
	parkingLot := createParkingLot(totalNumberOfSlots)

	for {
		fmt.Println("\n Available commands are \n 1: park registrationNumber color \n 2: leave slotNumber \n 3: status\n 4: slotNumbersForCarsWithColour color \n 5: slotNumberForCarsWithRegistrationNumber registrationNumber")
		reader := bufio.NewReader(os.Stdin)
		inputCommand, _ := reader.ReadString('\n')
		inputCommand = strings.TrimSuffix(inputCommand, "\n")

		command := strings.Split(inputCommand, " ")
		fmt.Println(command[0])

		switch command[0] {
		case "park":
			registrationNumber := command[1]
			color := command[2]
			parkingLot.parkCar(Car{registerationNumber: registrationNumber, color: color})
		case "leave":
			slotNumber, err := strconv.Atoi(command[1])
			fmt.Println(err)
			fmt.Println(slotNumber)
			parkingLot.freeSlot(slotNumber)
		case "status":
			parkingLot.getParkingLotStatus()
		case "slotNumbersForCarsWithColour":
			parkingLot.searchCarByColor(command[1])
		case "slotNumberForCarsWithRegistrationNumber":
			parkingLot.searchCarByRegisterationNumber(command[1])
		default:
			fmt.Println("Wrong command")
		}
	}

	// parkingLot.parkCar(Car{registerationNumber: "1234", color: "Green"})
	// parkingLot.parkCar(Car{registerationNumber: "1274", color: "Green"})
	// parkingLot.freeSlot(1)
	// parkingLot.parkCar(Car{registerationNumber: "5345", color: "Green"})
	// parkingLot.searchCarByColor("Green")
	// parkingLot.searchCarByRegisterationNumber("5345")
	// parkingLot.getParkingLotStatus()
}
