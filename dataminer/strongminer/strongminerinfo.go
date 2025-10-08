package strongminer

import (
	datamainer "nilchan/FinalProject/dataminer"
	"nilchan/FinalProject/datauser"
	"time"
)

type StrongMainerStr struct {
	datamainer.MinerBase
	count int
}

// вопрос
func NewStrongMainer() *StrongMainerStr {
	return &StrongMainerStr{
		MinerBase: datamainer.NewMiner("StrongMainer", 60),
		count:     0,
	}

}

func (m *StrongMainerStr) Mine() (int, error) {
	time.Sleep(1 * time.Second)
	if errEnergy := m.SetEnergy(1); errEnergy != nil {
		return 0, errEnergy
	}

	coal, errCoal := m.SetCoal(10 + m.count)
	if errCoal != nil {
		return 0, errCoal
	}
	m.count += 3
	return coal, nil
}

func (m *StrongMainerStr) Payment(salary int) (int, error) {
	subtraction := 450
	if salary < subtraction {
		return 0, datauser.ErrorLackingBalace
	}
	return subtraction, nil
}
