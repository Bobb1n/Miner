package http

import (
	"encoding/json"
	"errors"
	"nilchan/FinalProject/internal/models"
	"nilchan/FinalProject/internal/service/dataminer/minimainer"
	"nilchan/FinalProject/internal/service/dataminer/normalmainer"
	"nilchan/FinalProject/internal/service/dataminer/strongminer"
	"time"
)

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}
	return string(b)
}

type MinerDTO struct {
	Class string
}

func (m MinerDTO) ValidateToMinerDTO() (models.Miner, error) {
	if m.Class == "" {
		return nil, errors.New("the class field is empty")
	}
	switch {
	case m.Class == "Mini Mainer":
		miner := minimainer.NewMiniMainer()
		return miner, nil

	case m.Class == "Normal Mainer":
		miner := normalmainer.NewNormalMainer()
		return miner, nil

	case m.Class == "Strong Mainer":
		miner := strongminer.NewStrongMainer()
		return miner, nil
	}
	return nil, errors.New("не правильное имя класса")
}

type UpgradeDTO struct {
	Name      string `json:"name"`
	Count     int    `json:"count"`
	WasBought bool   `json:"was_bought"`
}

func (m UpgradeDTO) ValidateToUpgradeDTO() (models.Upgrade, error) {
	if m.Name == "" {
		return nil, errors.New("the Name field is empty")
	}
	switch {
	case m.Name == "Кирка":
		upgrade := models.NewPickaxe()
		return upgrade, nil

	case m.Name == "Вентиляция":
		upgrade := models.NewVentilation()
		return upgrade, nil

	case m.Name == "Вагонетка":
		upgrade := models.NewTrolley()
		return upgrade, nil
	}
	return nil, errors.New("не правильное имя апргрейда")
}
