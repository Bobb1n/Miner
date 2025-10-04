package minimainer

import (
	"math/rand"
	datamainer "nilchan/FinalProject/dataminer"
)

type MiniMainerStr struct {
	miniMainer *datamainer.Miner
}

// вопрос
func NewMiniMainer() *MiniMainerStr {
	id := rand.Intn(1000)
	return &MiniMainerStr{
		miniMainer: &datamainer.Miner{
			Id:        id,
			Level:     1,
			Class:     "MiniMainer",
			Energy:    30,
			Totalcoal: 0,
			IsRunning: false,
		},
	}
}

func (m *MiniMainerStr) IsRunning() {
	m.miniMainer.IsRunning = true
}

func (m *MiniMainerStr) UnIsRunning() {
	m.miniMainer.IsRunning = false
}

func (m MiniMainerStr) GetId() int {
	return m.miniMainer.Id
}
