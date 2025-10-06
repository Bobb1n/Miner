package minimainer

import (
	datamainer "nilchan/FinalProject/dataminer"
	"nilchan/FinalProject/datauser"
	"time"
)

type MiniMainerStr struct {
	datamainer.MinerBase
}

// вопрос
func NewMiniMainer() *MiniMainerStr {
	return &MiniMainerStr{
		datamainer.NewMiner("Mini Mainer", 30),
	}

}

func (m *MiniMainerStr) Mine() (int, error) {

	time.Sleep(3 * time.Second)
	if errEnergy := m.SetEnergy(1); errEnergy != nil {
		return 0, errEnergy
	}
	coal, errCoal := m.SetCoal(1)
	if errCoal != nil {
		return 0, errCoal
	}
	return coal, nil
}
func (m *MiniMainerStr) Payment(salary int) error {
	if salary <= 5 {
		return datauser.ErrorLackingBalace
	}
	return nil
}

// func (m *MiniMainerStr)Info()MiniMainerStr{
// 	tmp := make(map[int]MiniMainerStr)
// 	for k,v := range
// }
