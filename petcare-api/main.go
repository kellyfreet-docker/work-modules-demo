package main

import (
	"fmt"

	v2 "github.com/example/petcare-common/v2"
	"github.com/google/go-cmp/cmp"
)

func main() {
	pet1 := v2.Pet{
		ID:         "PET123",
		Name:       "Max",
		Species:    "Dog",
		Age:        5,
		OwnerEmail: "owner@example.com",
	}

	pet2 := v2.Pet{
		ID:         "PET123",
		Name:       "Max",
		Species:    "Dog",
		Age:        6,
		OwnerEmail: "newowner@example.com",
	}

	fmt.Println("API Service (using common v2)")
	fmt.Printf("Pet: %s (%s)\n", pet1.Name, pet1.ID)

	diff := cmp.Diff(pet1, pet2)
	fmt.Printf("Difference between pets:\n%s", diff)

	apt := v2.Appointment{
		ID:       "APT001",
		PetID:    pet1.ID,
		Time:     "2024-03-20 14:00",
		Status:   "Scheduled",
		Notes:    "Annual checkup",
		Duration: 30,
	}
	fmt.Println(v2.FormatAppointmentInfo(apt))
}
