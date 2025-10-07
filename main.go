package main

import (
	"fmt"
	datamainer "nilchan/FinalProject/dataminer"
	"nilchan/FinalProject/dataminer/minimainer"
	"time"

	"github.com/k0kubun/pp"
)

func main() {

	fmt.Println("Start program")

	miniMainer1 := minimainer.NewMiniMainer()

	MinerLogic := datamainer.NewManagerMainer()

	MinerLogic.AddMiner(miniMainer1)

	MinerLogic.Run(50, miniMainer1)

	// NormMiner := normalmainer.NewNormalMainer()

	// MinerLogic.AddMiner(NormMiner)

	// MinerLogic.Run(50, NormMiner)

	// StrongMiner := strongminer.NewStrongMainer()

	// MinerLogic.AddMiner(StrongMiner)

	// MinerLogic.Run(500, StrongMiner)
	time.Sleep(10 * time.Second)
	MinerLogic.StopAll()

	// time.Sleep(2 * time.Second)
	InfoByGroup := MinerLogic.InfoByGroup("Mini Mainer")
	infoAll := MinerLogic.InfoAll()
	pp.Println(InfoByGroup)
	pp.Println(infoAll)
	fmt.Println("Время кончилось")

}
