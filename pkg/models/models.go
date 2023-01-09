package models

type User struct {
	ID                              int
	Username, Password, Email, Name string
}

type Account struct {
	ID          int
	Email, Name string
}
