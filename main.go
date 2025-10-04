package main

import (
	"context"
	"fmt"
	datamainer "nilchan/FinalProject/dataminer"
	"nilchan/FinalProject/dataminer/minimainer"
	"nilchan/FinalProject/dataminer/normalmainer"
	"time"

	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("Start program")
	ctx := context.Background()

	miner := minimainer.NewMiniMainer()
	minerList := minimainer.NewList()
	minerList.AddMiner(miner)

	moduleMiner := datamainer.NewMainer(minerList)

	moduleMiner.Run(ctx, 50, miner.GetId(), wg)

	minernormal := normalmainer.NewNormMainer()
	minerList2 := normalmainer.NewNormalMiner()

	minerList2.AddMiner(&minernormal)
	minerModule2 := datamainer.NewMainer(minerList2)
	minerList2.Run(ctx, wg, 50, minernormal.GetId())
	// time.Sleep(1 * time.Second)
	// minerList.Start(ctx, 50, miner2.GetId(), wg)
	// time.Sleep(7 * time.Second)
	// for i := 0; i <= 1000; i++ {
	// 	minerList.Stop(i)
	// }
	wg.Wait()
	time.Sleep(2 * time.Second)

	infonormal := minerList2.Info(2)
	mininfo := minerList.Info(1)
	info := moduleMiner.InfoAll()
	fmt.Println(info)
	fmt.Println(mininfo)

	fmt.Println(infonormal)
	fmt.Println("Finished")
}
