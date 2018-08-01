package abstract_factory

import (
	"errors"
	"fmt"
)

type CarFactory struct{}

func (f *CarFactory) Build(t int) (Vehicle, error) {
	switch t {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d is not present\n", t))
	}
}
