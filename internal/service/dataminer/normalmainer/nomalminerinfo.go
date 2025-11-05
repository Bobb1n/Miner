package normalmainer

import (
	"nilchan/FinalProject/internal/models"
	"nilchan/FinalProject/pkg/errors"
	"time"
)

type NormMainerStr struct {
	models.MinerBase
}

// вопрос
func NewNormalMainer() *NormMainerStr {
	return &NormMainerStr{
		*models.NewMiner("Normal Mainer", 45),
	}

}

func (m *NormMainerStr) Mine() (int, error) {
	time.Sleep(2 * time.Second)
	if errEnergy := m.SetEnergy(1); errEnergy != nil {
		return 0, errEnergy
	}
	coal, errCoal := m.SetCoal(3)
	if errCoal != nil {
		return 0, errCoal
	}
	return coal, nil
}

func (m *NormMainerStr) Payment(salary int) (int, error) {
	subtraction := 50
	if salary < subtraction {
		return 0, errors.ErrorLackingBalace
	}
	return subtraction, nil
}
