package types

// Person: class
type Person struct {
	PersonId  string
	Firstname string
	Lastname  string
	Gender    string
}

// Employee: class
type Employee struct {
	EmployeeID       string
	P                Person
	EmployeePosition string
	EmployeeEmail    string
	EmployeeContact  string
	StartDate        string
	EndDate          string
}

// Enrollment: class
type Enrollment struct {
	EnrollmentId string
	P            Person
	Pr           Program
	EnrollDate   string
	UnenrollDate string
}
