package abstract_factory

import (
	"errors"
	"fmt"
)

type MotorbikeFactory struct{}

func (f *MotorbikeFactory) Build(t int) (Vehicle, error) {
	switch t {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d is not present\n", t))
	}
}
