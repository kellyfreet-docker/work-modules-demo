package v2

import "fmt"

func FormatAppointmentInfo(apt Appointment) string {
	return fmt.Sprintf("Appointment %s for pet %s at %s (Duration: %d mins)\nNotes: %s",
		apt.ID, apt.PetID, apt.Time, apt.Duration, apt.Notes)
}
