package normalmainer

import (
	datamainer "nilchan/FinalProject/dataminer"
	"nilchan/FinalProject/datauser"
	"time"
)

type NormMainerStr struct {
	datamainer.MinerBase
}

// вопрос
func NewNormalMainer() *NormMainerStr {
	return &NormMainerStr{
		*datamainer.NewMiner("Normal Mainer", 45),
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
		return 0, datauser.ErrorLackingBalace
	}
	return subtraction, nil
}
