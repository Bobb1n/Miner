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
	time         time.Time
	minerdata    map[string]int
}

func NewUser(name string) *User {
	user := &User{
		name:      name,
		upgrade:   []Upgrade{},
		time:      time.Now(),
		minerdata: make(map[string]int),
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
	Name      string           `json:"name"`
	Balance   int              `json:"balance"`
	Upgrades  map[string]int   `json:"upgrades"`
	Time      float64          `json:"game_time_seconds"`
	MinerData []map[string]int `json:"all miner in game"`
}

func (u *User) UserInfo() UserInfo {
	return UserInfo{
		Name:      u.name,
		Balance:   u.GetBalance(),
		Upgrades:  u.GetUpgradeStats(),
		Time:      time.Since(u.time).Seconds(),
		MinerData: u.minerTotalsSlice(),
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
func (u *User) AddMinerStats(class string, count int) {
	u.mtx.Lock()
	defer u.mtx.Unlock()
	u.minerdata[class] += count
}

func (u *User) minerTotalsSlice() []map[string]int {
	u.mtx.Lock()
	defer u.mtx.Unlock()
	result := make([]map[string]int, 0, len(u.minerdata))
	for class, count := range u.minerdata {
		result = append(result, map[string]int{class: count})
	}
	return result
}
