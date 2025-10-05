package normalmainer

import (
	"math/rand"
)

type NormMainer struct {
	id        int
	class     string
	energy    int
	totalcoal int
	isRunning bool
}

func NewNormMainer() *NormMainer {
	id := rand.Intn(1000)
	return &NormMainer{

		id:        id,
		class:     "Norm Miner",
		energy:    45,
		totalcoal: 0,
		isRunning: false,
	}
}

func (n *NormMainer) IsRunning() {
	n.isRunning = true
}
func (n *NormMainer) UnRunning() {
	n.isRunning = false
}
func (n *NormMainer) GetId() int {
	return n.id
}
