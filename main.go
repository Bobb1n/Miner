package main

import (
	"context"
	"fmt"
	datamainer "nilchan/FinalProject/dataminer"
	"nilchan/FinalProject/dataminer/minimainer"

	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("Start program")
	ctx := context.Background()

	miner := minimainer.NewMiniMainer()
	minerList := minimainer.NewList()
	minerList.AddMiner(&miner)

	datamainer.NewMainer(minerList)

	minerList.Start(ctx, 50, miner.GetId(), wg)
	// time.Sleep(1 * time.Second)
	// minerList.Start(ctx, 50, miner2.GetId(), wg)
	// time.Sleep(7 * time.Second)
	// for i := 0; i <= 1000; i++ {
	// 	minerList.Stop(i)
	// }
	time.Sleep(2 * time.Second)
	wg.Wait()
	info := minerList.InfoAll()
	fmt.Println(info)
	fmt.Println("Finished")
}
