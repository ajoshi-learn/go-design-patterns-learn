package composite

import "testing"

func TestTree(t *testing.T) {
	root := Tree{0, &Tree{1, nil, nil}, &Tree{2, nil, nil}}
	if root.Left.LeafValue != 2 {
		t.Error("Left leaf value must be equal to 2")
	}
}
