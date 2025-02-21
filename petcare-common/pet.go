package common

type Pet struct {
	ID      string
	Name    string
	Species string
	Age     int
}

type Appointment struct {
	ID     string
	PetID  string
	Time   string
	Status string
} 