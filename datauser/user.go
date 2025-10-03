package datauser

type User struct {
	Name    string
	Balance int
	Energy  int
}

func (u *User) NewUser(name string, balance int, energy int) User {
	return User{
		Name:    name,
		Balance: balance,
		Energy:  energy,
	}
}
