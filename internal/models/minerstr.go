package models

import (
	"errors"
	"sync"
	"sync/atomic"
)

var minerIDCounter atomic.Int64

type Miner interface {
	Mine() (int, error)
	GetId() int
	SetStatusWork(status bool)
	Payment(subtraction int) (int, error)
	GetStatusWork() bool
	Stats() MinerStats
}

type MinerBase struct {
	id         int
	class      string
	energy     atomic.Int64
	totalcoal  atomic.Int64
	statuswork atomic.Bool
	salary     atomic.Int64
	mtx        sync.Mutex
}

func NewMiner(class string, energy int, salary int) *MinerBase {
	id := int(minerIDCounter.Add(1))
	user := &MinerBase{
		id:    id,
		class: class,
	}
	user.energy.Store(int64(energy))
	user.salary.Store(int64(salary))
	user.totalcoal.Store(0)
	return user
}

// структура под  api
type MinerStats struct {
	Id         int    `json:"id"`
	Class      string `json:"class"`
	Energy     int    `json:"energy"`
	TotalCoal  int    `json:"total_coal"`
	StatusWork bool   `json:"status_work"`
}

func (m *MinerBase) Stats() MinerStats {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	return MinerStats{
		Id:         m.id,
		Class:      m.class,
		Energy:     int(m.energy.Load()),
		TotalCoal:  int(m.totalcoal.Load()),
		StatusWork: m.statuswork.Load(),
	}
}

func (m *MinerBase) GetEnergy() int   { return int(m.energy.Load()) }
func (m *MinerBase) GetSalary() int   { return int(m.salary.Load()) }
func (m *MinerBase) GetClass() string { return m.class }

func (m *MinerBase) GetId() int          { return m.id }
func (m *MinerBase) GetTotalCoal() int   { return int(m.totalcoal.Load()) }
func (m *MinerBase) GetStatusWork() bool { return m.statuswork.Load() }

func (m *MinerBase) SetCoal(coal int) (int, error) {
	if coal < 0 {
		return 0, errors.New("значение не может быть орицательным")
	}
	m.totalcoal.Add(int64(coal))
	return coal, nil
}
func (m *MinerBase) SetEnergy(energy int) error {
	if energy < 0 {
		return errors.New("значение не может быть орицательным")
	}
	for {
		current := m.energy.Load()
		if current < int64(energy) {
			return errors.New("добыча невозмажна так как кончилась энергия")
		}
		if m.energy.CompareAndSwap(current, current-int64(energy)) {
			return nil
		}

	}
}

func (m *MinerBase) SetStatusWork(status bool) {
	m.statuswork.Store(status)
}

type MinerTypeInfo struct {
	Class       string `json:"class"`
	Cost        int    `json:"cost"`
	Energy      int    `json:"energy"`
	Description string `json:"вescription"`
}
