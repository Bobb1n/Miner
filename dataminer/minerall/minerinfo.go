package minerall

import (
	datamainer "nilchan/FinalProject/dataminer"
)

type ListInfo interface {
	Info() map[int]datamainer.Miner
}

type MinerMap struct {
	groups []ListInfo
}

func (m *MinerMap) AddGroup(l ListInfo) {
	m.groups = append(m.groups, l)
}

func (m *MinerMap) InfoAll() map[int]datamainer.Miner {
	result := make(map[int]datamainer.Miner)
	for _, group := range m.groups {
		info := group.Info()
		for k, v := range info {
			result[k] = v
		}
	}
	return result
}
