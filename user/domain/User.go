package domain

type User struct {
	Name string
}

func (u User) Greet() string {
	return "User with name: " + u.Name + " created"
}
