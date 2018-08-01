package abstract_factory

const (
	SportMotorbikeType  = iota
	CruiseMotorbikeType
)

type Motorbike interface {
	GetMotorbikeType() int
}
