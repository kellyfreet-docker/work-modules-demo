package main

import (
	"fmt"

	"github.com/example/petcare-common"
	"github.com/google/go-cmp/cmp"
)

func main() {
	apt1 := common.Appointment{
		ID:     "APT001",
		PetID:  "PET123",
		Time:   "2024-03-20 14:00",
		Status: "Scheduled",
	}

	apt2 := common.Appointment{
		ID:     "APT001",
		PetID:  "PET123",
		Time:   "2024-03-20 15:00",
		Status: "Rescheduled",
	}

	fmt.Println("Scheduler Service")
	fmt.Println(common.FormatAppointmentInfo(apt1))

	diff := cmp.Diff(apt1, apt2)
	fmt.Printf("Appointment changes:\n%s", diff)
}
