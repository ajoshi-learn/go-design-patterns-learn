package abstract_factory

import (
	"testing"
)

func TestMotorbikeFactory(t *testing.T) {
	motobikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbike, err := motobikeF.Build(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels\n", motorbike.NumWheels())

	castedMotorbike, ok := motorbike.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Casted motorbike has type %d\n", castedMotorbike.GetMotorbikeType())
}
