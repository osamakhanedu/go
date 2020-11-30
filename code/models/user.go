package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users []*User // purpose of using point of User is it will
	// allow us to pass data all round the app without copying it.
	nextID uint32 = 1
)
