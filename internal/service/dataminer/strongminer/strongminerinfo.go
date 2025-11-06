package strongminer

import (
	"nilchan/FinalProject/internal/models"
	"nilchan/FinalProject/pkg/errors"
	"sync/atomic"
	"time"
)

const (
	ClassName       = "Strong Mainer"
	TypeCost        = 450
	TypeEnergy      = 60
	TypeDescription = "Мощный майнер, добывает 10+ угля в секунду"
)

func GetTypeInfo() models.MinerTypeInfo {
	return models.MinerTypeInfo{
		Class:       ClassName,
		Cost:        TypeCost,
		Energy:      TypeEnergy,
		Description: TypeDescription,
	}
}

type StrongMainerStr struct {
	models.MinerBase
	count atomic.Int32
}

// вопрос
func NewStrongMainer() *StrongMainerStr {
	return &StrongMainerStr{
		MinerBase: *models.NewMiner(ClassName, TypeEnergy, TypeCost),
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

func (m *StrongMainerStr) Payment(salaryGet int) (int, error) {

	if salaryGet < TypeCost {
		return 0, errors.ErrorLackingBalace
	}
	return TypeCost, nil
}
