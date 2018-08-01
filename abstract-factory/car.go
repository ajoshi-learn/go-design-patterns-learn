package abstract_factory

const (
	LuxuryCarType = iota
	FamilyCarType
)

type Car interface {
	NumDoors() int
}
