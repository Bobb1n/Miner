package obrobotchiki

import (
	"nilchan/FinalProject/datauser"
	"time"
)

func MiniMaiener(transferMiniChan chan int, salary int) (error, chan<- int) {
	if ok := salary < 5; !ok {
		return datauser.ErrorLackingBalace, nil
	}
	energy := 30
	sum := 0
	for i := 1; i <= energy; i++ {
		time.Sleep(3 * time.Second)
		sum += i

	}
	transferMiniChan <- sum
	return nil, transferMiniChan
}

func NormalMaiener(transferMiniChan chan int, salary int) (error, chan<- int) {
	if ok := salary < 50; !ok {
		return datauser.ErrorLackingBalace, nil
	}
	energy := 45
	sum := 0
	for i := 1; i <= energy; i++ {
		time.Sleep(2 * time.Second)
		sum += 3
	}
	transferMiniChan <- sum
	return nil, transferMiniChan
}

func StrongMaiener(transferMiniChan chan int, salary int) (error, chan<- int) {
	if ok := salary < 450; !ok {
		return datauser.ErrorLackingBalace, nil
	}
	energy := 60
	sum := 0
	for i := range energy {
		time.Sleep(1 * time.Second)
		sum += 10 + i*3
	}
	transferMiniChan <- sum
	return nil, transferMiniChan
}
