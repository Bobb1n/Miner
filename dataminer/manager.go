package datamainer

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type ManagerMainer struct {
	info   map[int]Miner
	mtx    sync.Mutex
	cancel map[int]context.CancelFunc
	done   map[int]chan struct{}
}

func NewManagerMainer() *ManagerMainer {
	return &ManagerMainer{
		info:   make(map[int]Miner),
		cancel: make(map[int]context.CancelFunc),
		done:   make(map[int]chan struct{}),
	}
}

func (m *ManagerMainer) AddMiner(miner Miner) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.info[miner.GetId()] = miner

}
func (m *ManagerMainer) Run(salary int, miner Miner) error {

	ctx := context.Background()
	// wg := &sync.WaitGroup{}

	m.mtx.Lock()
	defer m.mtx.Unlock()

	if errSalary := miner.Payment(salary); errSalary != nil {
		return errSalary
	}

	miner, ok := m.info[miner.GetId()]

	if !ok {
		return fmt.Errorf("майнер с таким id для запуска не найден")
	}

	if miner.GetStatusWork() {
		return fmt.Errorf("майнинг уже запущен")
	}
	minerCtx, minercansel := context.WithCancel(ctx)

	done := make(chan struct{})
	m.done[miner.GetId()] = done
	m.cancel[miner.GetId()] = minercansel
	miner.SetStatusWork(true)
	go m.Start(minerCtx, miner)

	return nil
}

func (m *ManagerMainer) Start(ctx context.Context, miner Miner) {
	// defer close(done)
	transferCoal := make(chan int)

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

				coal, err := miner.Mine()
				if err != nil {
					fmt.Println(err)
					return
				}
				transferCoal <- coal

			}
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		total := 0
		for coal := range transferCoal {
			total += coal
			log.Printf("Добыто: %d угля, Всего: %d \n", coal, total)
		}

	}()

	go func() {
		wg.Wait()
		m.mtx.Lock()
		miner.SetStatusWork(false)
		m.mtx.Unlock()
		fmt.Printf("Майнер %d автоматически остановлен\n", miner.GetId())

	}()

}

func (m *ManagerMainer) Stop(miner Miner) error {
	return nil
}

func (m *ManagerMainer) InfoById(id int) (MinerStats, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	miner, ok := m.info[id]
	if !ok {
		return MinerStats{}, fmt.Errorf("майнер с id %d не найден", id)
	}
	return miner.Stats(), nil
}

func (m *ManagerMainer) InfoByGroup(class string) []MinerStats {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	result := []MinerStats{}
	for _, miner := range m.info {
		stats := miner.Stats()
		if stats.Class == class {
			result = append(result, stats)
		}
	}
	return result
}
func (m *ManagerMainer) InfoAll() []MinerStats {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	result := []MinerStats{}
	for _, miner := range m.info {
		result = append(result, miner.Stats())
	}
	return result
}
