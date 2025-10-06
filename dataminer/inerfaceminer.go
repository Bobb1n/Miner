package datamainer

// import (
// 	"context"
// 	"sync"
// )

// type MinerMethods interface {
// 	Info() map[int]MinerBase
// 	Run(ctx context.Context, salary int, minerID int, wg *sync.WaitGroup) error

// 	// Stop(id int) error
// }

// type MinerModule struct {
// 	actionsModule MinerMethods
// 	minerInfo     map[int]*Miner
// }

// func NewMainer(MinerMethods MinerMethods) *MinerModule {
// 	return &MinerModule{
// 		actionsModule: MinerMethods,
// 		minerInfo:     make(map[int]*Miner),
// 	}
// }

// func (m MinerModule) Info() map[int]MinerBase {
// 	tmp := make(map[int]MinerBase, len(m.minerInfo))
// 	for k, v := range m.minerInfo {
// 		tmp[k] = *v
// 	}

// 	return tmp
// }

// // func (m *MinerModule) Info(level int) map[int]Miner {
// // 	return m.actionsModule.Info(level)

// func (m *MinerModule) Run(salary int, minerID int) error {

// 	ctx := context.Background()
// 	wg := &sync.WaitGroup{}
// 	return m.actionsModule.Run(ctx, salary, minerID, wg)
// }
// func Start() {

// }

// // func (m *MinerModule) Stop(id int) error {
// // 	return m.actionsModule.Stop(id)
// // }
