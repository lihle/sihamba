package types

// Login: class
type Login struct {
	LoginId   string
	E         Employee
	Username  string
	Password  string
	IsAdmin   int
	IsGeneral int
}

// Log: class
type Log struct {
	LogId       string
	L           Login
	LoginDate   string
	LoginStatus string
	LogoutDate  string
}
