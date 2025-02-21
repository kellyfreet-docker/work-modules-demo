package common

import "fmt"

func FormatAppointmentInfo(apt Appointment) string {
	return fmt.Sprintf("Appointment %s for pet %s at %s", apt.ID, apt.PetID, apt.Time)
} 