package minimainer

import (
	"context"
	"fmt"
	"log"
	"nilchan/FinalProject/datauser"
	"sync"
)

type List struct {
	info   map[int]*MiniMainerStr
	mtx    sync.Mutex
	cancel map[int]context.CancelFunc
	done   map[int]chan struct{}
}

// вопрос
func NewList() *List {
	return &List{
		info: make(map[int]*MiniMainerStr),
		done: make(map[int]chan struct{}),
	}
}

func (m *List) AddMiner(miner *MiniMainerStr) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.info[miner.GetId()] = miner
}

func (m *List) Start(ctx context.Context, salary int, minerID int, wg *sync.WaitGroup) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if salary <= 5 {
		return datauser.ErrorLackingBalace
	}

	miner, ok := m.info[minerID]

	if !ok {
		return fmt.Errorf("чот не так")
	}

	if miner.MiniMainerStr.IsRunning {
		return fmt.Errorf("майнинг уже запущен")
	}
	minerCtx, minercansel := context.WithCancel(ctx)
	if m.cancel == nil {
		m.cancel = make(map[int]context.CancelFunc)
	}
	done := make(chan struct{})
	m.done[minerID] = done
	m.cancel[minerID] = minercansel
	miner.IsRunning()
	m.info[miner.MiniMainerStr.Id] = miner

	go m.Run(minerCtx, minerID, done)

	return nil
}

func (m *List) Run(ctx context.Context, id int, done chan struct{}) {
	defer close(done)
	transferCoal := make(chan int)
	coal := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {

		defer wg.Done()
		defer close(transferCoal)

		for {
			select {
			case <-ctx.Done():

				fmt.Println("Finish")
				return

			default:

				m.mtx.Lock()
				miner, ok := m.info[id]
				if !ok || miner.MiniMainerStr.Energy <= 0 {
					m.mtx.Unlock()
					log.Println("завершили работу")
					return
				}

				coal += 1
				transferCoal <- coal

				miner.MiniMainerStr.Totalcoal = coal
				miner.MiniMainerStr.Energy = miner.MiniMainerStr.Energy - 1

				m.info[id] = miner
				m.mtx.Unlock()
				// time.Sleep(1 * time.Second)

			}
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for coal := range transferCoal {
			log.Printf("Обработано: %d угля \n", coal)
		}

	}()

	go func() {

		// defer close(transferCoal)
		m.mtx.Lock()
		miner := m.info[id]
		miner.UnIsRunning()
		m.info[id] = miner

		m.mtx.Unlock()
		wg.Wait()
		fmt.Printf("Майнер %d автоматически остановлен\n", id)

	}()

}

func (m *List) Stop(id int) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	miner, ok := m.info[id]
	if !ok {
		return fmt.Errorf("майнер не найден")

	}

	if cancel, exits := m.cancel[id]; exits && cancel != nil {
		cancel()
		delete(m.cancel, id)
	}

	miner.UnIsRunning()
	return nil
}

func (m *List) InfoAll() map[int]MiniMainerStr {
	tmp := make(map[int]MiniMainerStr, len(m.info))
	for k, v := range m.info {
		tmp[k] = *v
	}
	return tmp

}
