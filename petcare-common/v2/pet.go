package v2

type Pet struct {
	ID         string
	Name       string
	Species    string
	Age        int
	OwnerEmail string
}

type Appointment struct {
	ID       string
	PetID    string
	Time     string
	Status   string
	Notes    string
	Duration int // in minutes
}
