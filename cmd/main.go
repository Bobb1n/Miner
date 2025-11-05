package main

import (
	"fmt"
	"nilchan/FinalProject/internal/models"
	datamainer "nilchan/FinalProject/internal/service"
	"nilchan/FinalProject/internal/service/dataminer/minimainer"

	"time"

	"github.com/k0kubun/pp"
)

func main() {

	fmt.Println("Start program")

	user := models.NewUser("lala")
	pp.Println(user)
	user2 := models.UserInfo{
		Name:     "lalal",
		Balance:  100000,
		Upgrades: make(map[string]int),
	}
	pp.Println(user2)

	miniMainer1 := minimainer.NewMiniMainer()
	miniMainer2 := minimainer.NewMiniMainer()

	MinerLogic := datamainer.NewManagerMainer(user)

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
