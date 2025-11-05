package datamainer

import (
	"context"
	"fmt"
	"log"

	"nilchan/FinalProject/internal/models"

	"sync"
)

type ManagerMainer struct {
	info   map[int]models.Miner
	mtx    sync.Mutex
	cancel map[int]context.CancelFunc
	done   map[int]chan struct{}
	user   *models.User
}

func NewManagerMainer(user *models.User) *ManagerMainer {
	return &ManagerMainer{
		info:   make(map[int]models.Miner),
		cancel: make(map[int]context.CancelFunc),
		done:   make(map[int]chan struct{}),
		user:   user,
	}
}

func (m *ManagerMainer) AddMiner(miner models.Miner) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.info[miner.GetId()] = miner

}
func (m *ManagerMainer) Run(miner models.Miner) error {

	ctx := context.Background()
	wg := &sync.WaitGroup{}

	m.mtx.Lock()
	defer m.mtx.Unlock()

	subtraction, errSalary := miner.Payment(m.user.GetBalance())
	if errSalary != nil {
		return errSalary
	}
	m.user.SubtractBalance(subtraction)

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

	transferChan := m.Start(minerCtx, miner)
	//нужна ли тут овобще wg
	wg.Add(1)
	go func() {
		defer wg.Done()

		for coal := range transferChan {

			m.user.AddBalance(coal)
		}

	}()
	go func() {
		wg.Wait()
	}()

	return nil
}

func (m *ManagerMainer) Start(ctx context.Context, miner models.Miner) <-chan int {
	// defer close(done)

	transferCoal := make(chan int)
	total := 0

	wg := &sync.WaitGroup{}

	type MineResult struct {
		coal int
		err  error
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(transferCoal)

		for {
			resultChan := make(chan MineResult, 1)

			go func() {
				coal, err := miner.Mine()
				resultChan <- MineResult{coal: coal, err: err}
			}()

			select {
			case <-ctx.Done():

				fmt.Println("Майнинг остановлен по запросу")
				return

			case result := <-resultChan:
				if result.err != nil {
					fmt.Println("Ошибка майнинга:", result.err)
					return
				}

				select {
				case <-ctx.Done():

					return
				case transferCoal <- result.coal:
					total += result.coal
					log.Printf("Добыто: %d угля.Всего %d \n", result.coal, total)

				}
			}
		}
	}()

	go func() {
		wg.Wait()
		miner.SetStatusWork(false)
		log.Printf("Майнер %d автоматически остановлен\n", miner.GetId())

	}()
	return transferCoal

}

func (m *ManagerMainer) InfoById(id int) (models.MinerStats, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	miner, ok := m.info[id]
	if !ok {
		return models.MinerStats{}, fmt.Errorf("майнер с id %d не найден", id)
	}
	return miner.Stats(), nil
}

func (m *ManagerMainer) InfoByGroup(class string) []models.MinerStats {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	result := []models.MinerStats{}
	for _, miner := range m.info {
		stats := miner.Stats()
		if stats.Class == class {
			result = append(result, stats)
		}
	}
	return result
}
func (m *ManagerMainer) InfoAll() []models.MinerStats {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	result := []models.MinerStats{}
	for _, miner := range m.info {
		result = append(result, miner.Stats())
	}
	return result
}

// Стоп конкртенного манйера
func (m *ManagerMainer) Stop(id int) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	_, ok := m.info[id]
	if !ok {
		return fmt.Errorf("майнер с таким id не найден")
	}
	canсel, ok := m.cancel[id]
	if !ok {
		return fmt.Errorf("невозможно остановить майнер")
	}

	canсel()

	delete(m.cancel, id)

	return nil
}

func (m *ManagerMainer) StopGame() {
	m.mtx.Lock()

	cancels := make([]context.CancelFunc, 0, len(m.cancel))
	for _, cancel := range m.cancel {
		cancels = append(cancels, cancel)
	}

	m.cancel = make(map[int]context.CancelFunc)
	m.mtx.Unlock()

	for _, cancel := range cancels {
		cancel()
	}
	m.user.StopPassivIncome()

	log.Println("Сигнал остановки отправлен всем майнерам")
}
