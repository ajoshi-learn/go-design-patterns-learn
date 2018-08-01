package builder

import "testing"

func TestBuilderPattern(t *testing.T) {

	carBuilder := &CarBuilder{}

	car := carBuilder.SetWheels().SetSeats().SetStructure().Build()

	if car.Wheels != 4 {
		t.Errorf("Wheel on car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure of a car must be 'Car' and it was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Wheel on car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}

	bike := bikeBuilder.SetWheels().SetSeats().SetStructure().Build()

	if bike.Wheels != 2 {
		t.Errorf("Wheel on bike must be 2 and they were %d\n", bike.Wheels)
	}

	if bike.Structure != "Bike" {
		t.Errorf("Structure of a bike must be 'Bike' and it was %s\n", bike.Structure)
	}
}
