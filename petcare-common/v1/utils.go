package v1

import "fmt"

func FormatAppointmentInfo(apt Appointment) string {
	return fmt.Sprintf("Appointment %s for pet %s at %s", apt.ID, apt.PetID, apt.Time)
}
