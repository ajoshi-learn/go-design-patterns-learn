package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Cache is nil")
	}
	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	t.Log(&item1)
	t.Log(&whitePrototype)
	if item1 == whitePrototype {
		t.Error("Item can not be equal to prototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type assertion failed")
	}
	shirt1.SKU = "abbcc"
	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}
	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type assertion failed")
	}
	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU's must be different")
	}
	if shirt1 == shirt2 {
		t.Error("Shirts must be different")
	}
}
