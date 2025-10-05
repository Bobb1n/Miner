package normalmainer

import (
	"context"
	"errors"
	"fmt"
	datamainer "nilchan/FinalProject/dataminer"
	"nilchan/FinalProject/datauser"
	"sync"
	"time"
)

type List struct {
	info   map[int]*NormMainer
	mtx    sync.Mutex
	cancel map[int]context.CancelFunc
	done   map[int]chan struct{}
}

func NewNormalMiner() *List {
	return &List{
		info:   make(map[int]*NormMainer),
		cancel: make(map[int]context.CancelFunc),
		done:   make(map[int]chan struct{}),
	}
}

func (m *List) AddMiner(miner *NormMainer) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.info[miner.id] = miner
}

func (m *List) Run(ctx context.Context, salary int, id int, wg *sync.WaitGroup) error {
	if salary < 50 {
		return datauser.ErrorLackingBalace
	}
	miner, ok := m.info[id]
	if !ok {
		return errors.New("такого майнера нет")
	}
	if miner.isRunning {
		return errors.New("уже запущен майнер")
	}
	normcontext, cancelnorm := context.WithCancel(ctx)
	miner.IsRunning()
	//need test
	done := make(chan struct{})
	m.cancel[id] = cancelnorm
	go m.Start(normcontext, id, done)
	return nil
}
func (m *List) Start(ctx context.Context, id int, done chan struct{}) {
	defer close(done)
	miner, _ := m.info[id]
	transferPoint := make(chan int)
	wg := &sync.WaitGroup{}
	coal := 0
	wg.Add(1)
	go func() {

		defer close(transferPoint)
		defer wg.Done()

		select {
		case <-ctx.Done():
			fmt.Println("Finish")
			return
		default:
			for {

				m.mtx.Lock()

				if miner.energy <= 0 {
					fmt.Println("Энергия кончилась")
					return
				}
				miner.energy--
				coal = coal + 3
				miner.totalcoal = coal
				transferPoint <- coal
				m.mtx.Unlock()
				time.Sleep(100 * time.Millisecond)
			}

		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range transferPoint {
			fmt.Println("Обработанно:", v, "угля")
		}
	}()
	go func() {
		wg.Wait()
		miner.UnRunning()
	}()
}

// func (m *List) Info(level int) map[int]datamainer.Miner {
// 	tmp := make(map[int]datamainer.Miner, len(m.info))
// 	for k, v := range m.info {
// 		if v.normMiner.Level == level {
// 			tmp[k] = *v.normMiner
// 		}
// 	}
// 	return tmp
// }

func (m *List) Info() map[int]datamainer.Miner {

	miners := make([]*NormMainer, 0, len(m.info))
	for _, v := range m.info {
		miners = append(miners, v)
	}

	tmp := make(map[int]datamainer.Miner, len(m.info))
	for k, v := range miners {
		tmp[k] = datamainer.Miner{
			Id:        v.id,
			Class:     v.class,
			Energy:    v.energy,
			Totalcoal: v.totalcoal,
			IsRunning: v.isRunning,
		}
	}
	return tmp
}
