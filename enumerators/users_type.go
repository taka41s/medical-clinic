type UserType int

const (
	// UserTypePatient represents a user type of patient
	UserTypePatient UserType = iota
	// UserTypeDoctor represents a user type of doctor
	UserTypeDoctor
)

// String returns the string representation of UserType
func (ut UserType) String() string {
	switch ut {
	case UserTypePatient:
		return "Patient"
	case UserTypeDoctor:
		return "Doctor"
	default:
		return "Unknown"
	}
}
