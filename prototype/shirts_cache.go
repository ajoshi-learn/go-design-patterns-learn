package prototype

import "errors"

const (
	White = iota
	Black
	Blue
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

type ShirtsCache struct {
}

func (s *ShirtsCache) GetClone(t int) (ItemInfoGetter, error) {
	switch t {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("shirt model not recognized")
	}
}

func GetShirtsCloner() ShirtCloner {
	return new(ShirtsCache)
}

var whitePrototype = &Shirt{15.00, "empty", White}
var blackPrototype = &Shirt{15.00, "empty", Black}
var bluePrototype = &Shirt{15.00, "empty", Blue}
