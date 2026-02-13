package minimainer

import (
	"nilchan/FinalProject/internal/models"
	"nilchan/FinalProject/pkg/errors"
	"time"
)

const (
	TypeClass       = "Mini Mainer"
	TypeEnergy      = 30
	TypeSalary      = 5
	TypeDescription = "Базовый майнер, добывает 1 угля в секунду"
)

type MiniMainerStr struct {
	models.MinerBase
}

func NewMiniMainer() *MiniMainerStr {
	return &MiniMainerStr{
		*models.NewMiner(TypeClass, TypeEnergy, TypeSalary),
	}
}

func GetTypeInfo() models.MinerTypeInfo {
	return models.MinerTypeInfo{
		Class:       TypeClass,
		Cost:        TypeSalary,
		Energy:      TypeEnergy,
		Description: TypeDescription,
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

func (m *MiniMainerStr) Payment(salaryGet int) (int, error) {
	if salaryGet < TypeSalary {
		return 0, errors.ErrorLackingBalace
	}
	return TypeSalary, nil
}
