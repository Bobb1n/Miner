package datauser

import (
	"context"
	"fmt"
	"nilchan/FinalProject/upgrade"
	"sync/atomic"
	"time"
)

type User struct {
	name         string
	balance      atomic.Int64 //TotalCoal
	upgrade      upgrade.Upgrade
	passivIncome context.CancelFunc
}

func NewUser(name string) *User {
	user := &User{
		name: name,
	}
	user.balance.Store(100)
	ctx, cansel := context.WithCancel(context.Background())
	user.passivIncome = cansel
	go user.PassivIncome(ctx)
	return user
}

func (u *User) GetBalance() int { return int(u.balance.Load()) }
func (u *User) GetName() string { return u.name }
func (u *User) AddBalance(amount int) {
	u.balance.Add(int64(amount))
}
func (u *User) SubtractBalance(amount int) error {
	result := u.balance.Load()
	if result < int64(amount) {
		return ErrorLackingBalace
	}
	u.balance.Add(-int64(amount))
	return nil
}
func (u *User) InfoUser() {

	fmt.Printf(`
User Information:
  Name:    %s
  Balance: %d
`, u.name, u.balance.Load())
}

func (u *User) Upgrade(updrade upgrade.Upgrade) {
	u.upgrade.BuyUpgrade()
}

func (m *User) PassivIncome(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Закончили")

		case <-ticker.C:
			m.AddBalance(1)

		}

	}
}
