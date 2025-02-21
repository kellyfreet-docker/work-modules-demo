package main

import (
	"fmt"

	"github.com/example/petcare-common"
	"github.com/google/go-cmp/cmp"
)

func main() {
	pet1 := common.Pet{
		ID:      "PET123",
		Name:    "Max",
		Species: "Dog",
		Age:     5,
	}

	pet2 := common.Pet{
		ID:      "PET123",
		Name:    "Max",
		Species: "Dog",
		Age:     6,
	}

	fmt.Println("API Service")
	fmt.Printf("Pet: %s (%s)\n", pet1.Name, pet1.ID)

	diff := cmp.Diff(pet1, pet2)
	fmt.Printf("Difference between pets:\n%s", diff)

	apt := common.Appointment{
		ID:     "APT001",
		PetID:  pet1.ID,
		Time:   "2024-03-20 14:00",
		Status: "Scheduled",
	}
	fmt.Println(common.FormatAppointmentInfo(apt))
}
