package strongminer

import (
	"nilchan/FinalProject/internal/models"
	"nilchan/FinalProject/pkg/errors"
	"sync/atomic"
	"time"
)

type StrongMainerStr struct {
	models.MinerBase
	count atomic.Int32
}

// вопрос
func NewStrongMainer() *StrongMainerStr {
	return &StrongMainerStr{
		MinerBase: *models.NewMiner("StrongMainer", 60),
	}

}

func (m *StrongMainerStr) Mine() (int, error) {
	time.Sleep(1 * time.Second)
	currentAddition := int(m.count.Load())
	if errEnergy := m.SetEnergy(1); errEnergy != nil {
		return 0, errEnergy
	}

	coal, errCoal := m.SetCoal(10 + currentAddition)
	if errCoal != nil {
		return 0, errCoal
	}
	m.count.Add(3)
	return coal, nil
}

func (m *StrongMainerStr) Payment(salary int) (int, error) {
	subtraction := 450
	if salary < subtraction {
		return 0, errors.ErrorLackingBalace
	}
	return subtraction, nil
}
