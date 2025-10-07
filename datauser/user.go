package datauser

type User struct {
	Name    string
	Balance int //TotalCoal
	Energy  int
}

func (u *User) NewUser(name string, balance int, energy int) User {
	return User{
		Name:    name,
		Balance: balance,
		Energy:  energy,
	}
}

// func (u *User) BalanceAdd() {

// 	go func() {
// 		defer wg.Done()
// 		for v := range coalTransferPoint {
// 			coal.Add(int64(v))
// 		}
// 	}()
// }
