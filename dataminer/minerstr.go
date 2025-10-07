package datamainer

import (
	"errors"
	"sync"
	"sync/atomic"
)

var minerIDCounter int

type Miner interface {
	Mine() (int, error)
	GetId() int
	SetStatusWork(status bool)
	Payment(salary int) error
	GetStatusWork() bool
	Stats() MinerStats
}

type MinerBase struct {
	id         int
	class      string
	energy     int
	totalcoal  int
	statuswork atomic.Bool
	mtx        sync.Mutex
}

func NewMiner(class string, energy int) MinerBase {
	minerIDCounter++
	return MinerBase{
		id:        minerIDCounter,
		class:     class,
		energy:    energy,
		totalcoal: 0,
	}
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
		Energy:     m.energy,
		TotalCoal:  m.totalcoal,
		StatusWork: m.statuswork.Load(),
	}
}

func (m MinerBase) GetEnergy() int { return m.energy }

func (m MinerBase) GetId() int          { return m.id }
func (m MinerBase) GetTotalCoal() int   { return m.totalcoal }
func (m MinerBase) GetStatusWork() bool { return m.statuswork.Load() }

func (m *MinerBase) SetCoal(coal int) (int, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if coal < 0 {
		return 0, errors.New("значение не может быть орицательным")
	}
	m.totalcoal += coal
	return coal, nil
}
func (m *MinerBase) SetEnergy(energy int) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if m.energy <= 0 {
		return errors.New("добыча невозмажна так как кончилась энергия")
	}
	if energy < 0 {
		return errors.New("значение не может быть орицательным")
	}
	m.energy -= energy
	return nil
}

// реализовать атомик тут
func (m *MinerBase) SetStatusWork(status bool) {
	m.statuswork.Store(status)

}
