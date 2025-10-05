package datamainer

import (
	"context"
	"sync"
)

type MinerMethods interface {
	Info() map[int]Miner
	Run(ctx context.Context, salary int, minerID int, wg *sync.WaitGroup) error
	// Info(level int) map[int]Miner
	// Stop(id int) error
}

type MinerModule struct {
	actionsModule MinerMethods
	minerInfo     map[int]*Miner
}

func NewMainer(MinerMethods MinerMethods) *MinerModule {
	return &MinerModule{
		actionsModule: MinerMethods,
		minerInfo:     make(map[int]*Miner),
	}
}

func (m *MinerModule) Info() map[int]Miner {

	return m.actionsModule.Info()
}

// func (m *MinerModule) Info(level int) map[int]Miner {
// 	return m.actionsModule.Info(level)

func (m *MinerModule) Run(ctx context.Context, salary int, minerID int, wg *sync.WaitGroup) error {
	return m.actionsModule.Run(ctx, salary, minerID, wg)
}

// func (m *MinerModule) Stop(id int) error {
// 	return m.actionsModule.Stop(id)
// }
