package pidroponics

type SamplePoint int

const (
	SUMP SamplePoint = iota
	INLET
	OUTLET
)

func GetSamplePoint(i int) SamplePoint {
	switch i {
		case 0:
			return SUMP
		case 1:
			return INLET
		case 2:
			return OUTLET
	}
	return -1
}

func (sp SamplePoint) String() string {
	return [...]string{"Sump", "Inlet", "Outlet"}[sp]
}
