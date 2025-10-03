package normalmainer

import "math/rand"

type NormMainer struct {
	class     string
	energy    int
	totalcoal int
	isRunning bool
	id        int
}

func NewNormMainer() NormMainer {
	id := rand.Intn(1000)
	return NormMainer{
		class:     "NormMainer",
		energy:    45,
		totalcoal: 0,
		isRunning: false,
		id:        id,
	}
}

func (n *NormMainer) IsRunning() {
	n.isRunning = true
}
func (n *NormMainer) UnRunning() {
	n.isRunning = false
}
