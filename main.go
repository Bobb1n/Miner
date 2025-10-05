package main

import (
	"context"
	"fmt"
	datamainer "nilchan/FinalProject/dataminer"
	"nilchan/FinalProject/dataminer/minerall"
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

	miner2 := minimainer.NewMiniMainer()
	minerList1 := minimainer.NewList()
	minerList1.AddMiner(miner2)

	moduleMiner1 := datamainer.NewMainer(minerList1)

	moduleMiner1.Run(ctx, 50, miner.GetId(), wg)

	minernormal := normalmainer.NewNormMainer()
	minerList2 := normalmainer.NewNormalMiner()

	minerList2.AddMiner(minernormal)
	minerModule2 := datamainer.NewMainer(minerList2)
	minerModule2.Run(ctx, 50, minernormal.GetId(), wg)

	minernormal2 := normalmainer.NewNormMainer()
	minerList3 := normalmainer.NewNormalMiner()

	minerList3.AddMiner(minernormal2)
	minerModule3 := datamainer.NewMainer(minerList2)
	minerModule3.Run(ctx, 50, minernormal2.GetId(), wg)

	// time.Sleep(1 * time.Second)
	// minerList.Start(ctx, 50, miner2.GetId(), wg)
	// time.Sleep(7 * time.Second)
	// for i := 0; i <= 1000; i++ {
	// 	minerList.Stop(i)
	// }
	wg.Wait()
	time.Sleep(20 * time.Second)

	infonormal := moduleMiner.Info()

	infoAll := &minerall.MinerMap{}
	infoAll.AddGroup(minerList)

	infoAll.AddGroup(minerList2)

	allinfo := infoAll.InfoAll()
	fmt.Println(allinfo)
	fmt.Println(infonormal)
	fmt.Println("Finished")
}
