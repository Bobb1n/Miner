package datamainer

import (
	"context"
	"fmt"
	"sync"
)

type MinerMethods interface {
	InfoAll() map[int]Miner
	Run(ctx context.Context, id int, done chan struct{})
}

type MinerModule struct {
	ActionsModule MinerMethods
	Info          map[int]*Miner
	Cancel        map[int]context.CancelFunc
	Done          map[int]chan struct{}
	Mtx           sync.Mutex
}

func NewMainer(MinerMethods MinerMethods) MinerModule {
	return MinerModule{
		ActionsModule: MinerMethods,
		Info:          make(map[int]*Miner),
		Cancel:        make(map[int]context.CancelFunc),
		Done:          make(map[int]chan struct{}),
	}
}

func (m MinerModule) InfoAll() map[int]Miner {
	tmp := make(map[int]Miner, len(m.Info))
	for k, v := range m.Info {
		tmp[k] = *v
	}

	fmt.Println("Вывод актуальнойц информации о всех рабочих")
	return tmp

}
