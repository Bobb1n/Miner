package minimainer

import (
	"math/rand"
	datamainer "nilchan/FinalProject/dataminer"
)

type MiniMainerStr struct {
	MiniMainerStr datamainer.Miner
}

// вопрос
func NewMiniMainer() MiniMainerStr {
	id := rand.Intn(1000)
	return MiniMainerStr{
		MiniMainerStr: datamainer.Miner{
			Id:        id,
			Class:     "MiniMainer",
			Energy:    30,
			Totalcoal: 0,
			IsRunning: false,
		},
	}
}

func (m *MiniMainerStr) IsRunning() {
	m.MiniMainerStr.IsRunning = true
}

func (m *MiniMainerStr) UnIsRunning() {
	m.MiniMainerStr.IsRunning = false
}

func (m MiniMainerStr) GetId() int {
	return m.MiniMainerStr.Id
}
