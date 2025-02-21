package main

import (
	"fmt"

	"github.com/example/petcare-common/v1"
	"github.com/google/go-cmp/cmp"
)

func main() {
	apt1 := v1.Appointment{
		ID:     "APT001",
		PetID:  "PET123",
		Time:   "2024-03-20 14:00",
		Status: "Scheduled",
	}

	apt2 := v1.Appointment{
		ID:     "APT001",
		PetID:  "PET123",
		Time:   "2024-03-20 15:00",
		Status: "Rescheduled",
	}

	fmt.Println("Scheduler Service (using common v1)")
	fmt.Println(v1.FormatAppointmentInfo(apt1))

	diff := cmp.Diff(apt1, apt2)
	fmt.Printf("Appointment changes:\n%s", diff)
}
