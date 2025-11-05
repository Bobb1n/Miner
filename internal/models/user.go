package models

import (
	"context"
	"fmt"
	"nilchan/FinalProject/pkg/errors"
	"sync"
	"sync/atomic"
	"time"
)

type User struct {
	name         string
	balance      atomic.Int64 //TotalCoal
	upgrade      []Upgrade
	passivIncome context.CancelFunc
	mtx          sync.Mutex
}

func NewUser(name string) *User {
	user := &User{
		name:    name,
		upgrade: []Upgrade{},
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
	for {
		result := u.balance.Load()
		if result < int64(amount) {
			return errors.ErrorLackingBalace
		}

		if u.balance.CompareAndSwap(result, result-int64(amount)) {
			return nil
		}
	}

}

type UserInfo struct {
	Name     string         `json:"name"`
	Balance  int            `json:"balance"`
	Upgrades map[string]int `json:"upgrades"`
}

func (u *User) UserInfo() UserInfo {
	return UserInfo{
		Name:     u.name,
		Balance:  u.GetBalance(),
		Upgrades: u.GetUpgradeStats(),
	}
}

func (u *User) Upgrade(updrade Upgrade) error {
	if updrade == nil {
		return fmt.Errorf("апгрейд не может быть nil")
	}
	price := updrade.BuyUpgrade()
	err := u.SubtractBalance(price)
	if err != nil {
		return err
	}

	u.mtx.Lock()
	u.upgrade = append(u.upgrade, updrade)
	u.mtx.Unlock()

	return nil
}

//	func (u *User) GetInfoUpgrades() []upgrade.Upgrade {
//		u.mtx.Lock()
//		defer u.mtx.Unlock()
//		result := make([]upgrade.Upgrade, len(u.upgrade))
//		copy(result, u.upgrade)
//		return result
//	}

func (u *User) GetUpgradeStats() map[string]int {
	u.mtx.Lock()
	defer u.mtx.Unlock()

	stats := make(map[string]int)
	for _, upg := range u.upgrade {
		stats[upg.GetName()]++
	}

	return stats
}

func (m *User) PassivIncome(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Пассивный доход остановлен")
			return

		case <-ticker.C:
			m.AddBalance(1)
		}
	}
}

func (u *User) StopPassivIncome() {
	if u.passivIncome != nil {
		u.passivIncome()
	}

}
