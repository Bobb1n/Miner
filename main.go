package main

import (
	"fmt"
	datamainer "nilchan/FinalProject/dataminer"
	"nilchan/FinalProject/dataminer/minimainer"
	"nilchan/FinalProject/datauser"
	"time"

	"github.com/k0kubun/pp"
)

func main() {

	fmt.Println("Start program")

	NewUSer := datauser.NewUser("Vanya")

	miniMainer1 := minimainer.NewMiniMainer()
	miniMainer2 := minimainer.NewMiniMainer()

	MinerLogic := datamainer.NewManagerMainer(NewUSer)

	MinerLogic.AddMiner(miniMainer1)
	MinerLogic.AddMiner(miniMainer2)

	err := MinerLogic.Run(miniMainer1)
	fmt.Println(err)
	err = MinerLogic.Run(miniMainer2)
	fmt.Println(err)

	// NormMiner := normalmainer.NewNormalMainer()

	// MinerLogic.AddMiner(NormMiner)

	// MinerLogic.Run(50, NormMiner)

	// StrongMiner := strongminer.NewStrongMainer()

	// MinerLogic.AddMiner(StrongMiner)

	// MinerLogic.Run(500, StrongMiner)

	time.Sleep(33 * time.Second)
	// MinerLogic.StopAll()

	// time.Sleep(2 * time.Second)
	InfoByGroup := MinerLogic.InfoByGroup("Mini Mainer")
	infoAll := MinerLogic.InfoAll()
	pp.Println(InfoByGroup)
	pp.Println(infoAll)
	fmt.Println("Время кончилось")

}
