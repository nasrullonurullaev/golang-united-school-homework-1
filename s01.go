package structs

import "fmt"

type User struct {
	firstName string
	lastName  string
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.lastName, u.firstName)
}

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}

func New() *User {
	return &User{}
}

func ResetUser(u *User) {
	u.firstName = ""
	u.lastName = ""
}

func IsUser(input interface{}) bool {
	_, ok := input.(*User)
	return ok
}

func ProcessUser(ui UserInterface) string {
	u := New()
	u.SetFirstName("firstName")
	u.SetLastName("lastName")
	ui.SetFirstName(u.firstName)
	ui.SetLastName(u.lastName)
	return ui.FullName()
}
