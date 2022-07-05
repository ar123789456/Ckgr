package models

type User struct {
	ID       int
	Nick     string
	FullName map[Language]string
	Special  map[Language]string
	Session  string
	Password string
}
