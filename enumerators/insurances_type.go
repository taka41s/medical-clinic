type InsuranceType int

const (
	Aetna InsuranceType = iota
	Healthcare
	Cigna
)

func (ut InsuranceType) String() string {
	switch ut {
	case Aetna:
		return "Aetna"
	case Healthcare:
		return "Healthcare"
	case Cigna:
		return "Cigna"
	}
}
