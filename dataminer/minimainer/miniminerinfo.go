package minimainer

import (
	"math/rand"
)

type MiniMainerStr struct {
	id        int
	class     string
	energy    int
	totalcoal int
	isRunning bool
}

// вопрос
func NewMiniMainer() *MiniMainerStr {
	id := rand.Intn(1000)
	return &MiniMainerStr{
		id:        id,
		class:     "MiniMainer",
		energy:    30,
		totalcoal: 0,
		isRunning: false,
	}

}

func (m *MiniMainerStr) IsRunning() {
	m.isRunning = true
}

func (m *MiniMainerStr) UnIsRunning() {
	m.isRunning = false
}

func (m MiniMainerStr) GetId() int {
	return m.id
}
