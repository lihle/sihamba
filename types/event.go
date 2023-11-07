package types

// Event: class
type Event struct {
	EventId    string
	Pr         Program
	EventDate  string
	EventName  string
	EventVenue string
}

// Attendance: class
type Attendance struct {
	Ev             Event
	En             Enrollment
	NumberAttended int
	NumberEnrolled int
}
