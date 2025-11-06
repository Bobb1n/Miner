package normalmainer

import (
	"nilchan/FinalProject/internal/models"
	"nilchan/FinalProject/pkg/errors"
	"time"
)

const (
	ClassName       = "Normal Mainer"
	TypeCost        = 50
	TypeEnergy      = 45
	TypeDescription = "Средний майнер, добывает 3 угля за 2 секунды"
)

func GetTypeInfo() models.MinerTypeInfo {
	return models.MinerTypeInfo{
		Class:       ClassName,
		Cost:        TypeCost,
		Energy:      TypeEnergy,
		Description: TypeDescription,
	}
}

type NormMainerStr struct {
	models.MinerBase
}

// вопрос
func NewNormalMainer() *NormMainerStr {
	return &NormMainerStr{
		*models.NewMiner(ClassName, TypeEnergy, TypeCost),
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

func (m *NormMainerStr) Payment(salaryGet int) (int, error) {
	if salaryGet < TypeCost {
		return 0, errors.ErrorLackingBalace
	}
	return TypeCost, nil
}
