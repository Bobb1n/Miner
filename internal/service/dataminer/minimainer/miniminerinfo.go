package minimainer

import (
	"nilchan/FinalProject/internal/models"
	"nilchan/FinalProject/pkg/errors"
	"time"
)

type MiniMainerStr struct {
	models.MinerBase
}

// вопрос
func NewMiniMainer() *MiniMainerStr {
	return &MiniMainerStr{
		*models.NewMiner("Mini Mainer", 30),
	}

}

func (m *MiniMainerStr) Mine() (int, error) {

	time.Sleep(1 * time.Second)
	if errEnergy := m.SetEnergy(1); errEnergy != nil {
		return 0, errEnergy
	}
	coal, errCoal := m.SetCoal(1)
	if errCoal != nil {
		return 0, errCoal
	}
	return coal, nil
}
func (m *MiniMainerStr) Payment(salary int) (int, error) {
	subtraction := 5
	if salary < subtraction {
		return 0, errors.ErrorLackingBalace
	}
	return subtraction, nil
}

// func (m *MiniMainerStr)Info()MiniMainerStr{
// 	tmp := make(map[int]MiniMainerStr)
// 	for k,v := range
// }
