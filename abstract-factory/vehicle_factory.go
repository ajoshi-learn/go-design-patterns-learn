package abstract_factory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	Build(VehicleType int) (Vehicle, error)
}

type FactoryType int

const (
	CarFactoryType       FactoryType = iota
	MotorbikeFactoryType
)

func BuildFactory(t FactoryType) (VehicleFactory, error) {
	switch t {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d is not recognized", t))
	}
}
