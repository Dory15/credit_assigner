package models

type Investment struct {
	Investment int32 `json:"investment"`
}

type Credits struct {
	ThreeHundred int32 `json:"credit_type_300"`
	FiveHundred  int32 `json:"credit_type_500"`
	SevenHundred int32 `json:"credit_type_700"`
}

type CreditStatus struct {
	AssigmentSuccessful bool `json:"assigment_successful"`
}

type CreditStatistics struct {
	AssigmentsMade               int32   `json:"Total de asignaciones realizadas" db:"assignments_made"`
	SuccessfulAssignments        int32   `json:"Total de asignaciones exitosas" db:"successful_assignments"`
	FailedAssignments            int32   `json:"Total de asignaciones no exitosas" db:"failed_assignments"`
	AverageSuccessfulAssignments float32 `json:"Promedio de inversión exitosa" db:"average_successful_assignments"`
	AverageFailedAssignments     float32 `json:"Promedio de inversión no exitosa" db:"average_failed_assignments"`
}
