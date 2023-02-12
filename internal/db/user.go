package db

type User struct {
	Username  string
	Firstname string
	Lastname  string
	Age       int
}

type UserRepository interface {
}
