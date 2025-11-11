package http

import (
	"encoding/json"
	"net/http"
	"nilchan/FinalProject/internal/models"
	datamainer "nilchan/FinalProject/internal/service"
	"nilchan/FinalProject/internal/service/dataminer/minimainer"
	"nilchan/FinalProject/internal/service/dataminer/normalmainer"
	"nilchan/FinalProject/internal/service/dataminer/strongminer"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type HttpHandlers struct {
	MinerManager *datamainer.ManagerMainer
	User         *models.User
}

func NewHttpHandlers(minerManager *datamainer.ManagerMainer, user *models.User) *HttpHandlers {
	return &HttpHandlers{
		MinerManager: minerManager,
		User:         user,
	}
}

func sendJsonError(w http.ResponseWriter, code int, messange string) {
	errorDTO := ErrorDTO{
		Message: messange,
		Time:    time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorDTO)
}

func sendJsonSuccses(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
pattern: /api/user
method:  GET
info:    -

succeed:
  - status code:   200 Ok
  - response body: JSON represent get info

failed:
  - status code:   400, 500, ...
  - response body: JSON with error + time
*/
func (h *HttpHandlers) HandleGetInfoUser(w http.ResponseWriter, r *http.Request) {
	user := h.User.UserInfo()
	sendJsonSuccses(w, http.StatusOK, user)

}

/*
pattern: /api/user/stop
method:  POST
info:    _
succeed:
	- status code:   200 OK
	- response body: JSON represent change status game

failed:
	- status code:   400, 500, ...
	- response body: JSON with error + time
*/

func (h *HttpHandlers) HandleStopGame(w http.ResponseWriter, r *http.Request) {
	userInfo, _ := h.MinerManager.StopGame()
	sendJsonSuccses(w, http.StatusOK, userInfo)

}

/*
pattern: /api/miners/types
method:  GET
info:    -

succeed:
	- status code:   200 OK
	- response body: JSON represent get info

failed:
	- status code:   400, 500, ...
	- response body: JSON with error + time
*/

func (h *HttpHandlers) HandleInfoSalryMiner(w http.ResponseWriter, r *http.Request) {
	types := []models.MinerTypeInfo{
		minimainer.GetTypeInfo(),
		normalmainer.GetTypeInfo(),
		strongminer.GetTypeInfo(),
	}
	sendJsonSuccses(w, http.StatusOK, types)

}

/*
pattern: /api/miners
method:  POST
info:    JSON in HTTP request body

succeed:
	- status code:   201 OK
	- response body: JSON represent create new miner

failed:
	- status code:   400, 409, 500, ...
	- response body: JSON with error + time
*/

func (h *HttpHandlers) HandleAddNewMiner(w http.ResponseWriter, r *http.Request) {
	var minerDTO MinerDTO
	if err := json.NewDecoder(r.Body).Decode(&minerDTO); err != nil {
		sendJsonError(w, 400, err.Error())
		return
	}
	defer r.Body.Close()
	miner, err := minerDTO.ValidateToMinerDTO()
	if err != nil {
		sendJsonError(w, 400, err.Error())
		return
	}
	h.MinerManager.AddMiner(miner)
	if err := h.MinerManager.Run(miner); err != nil {
		sendJsonError(w, 500, err.Error())
		return
	}
	miner.Stats()
	sendJsonSuccses(w, 201, miner.Stats())
}

/*
pattern: /api/miners?status=true or ?class=Mini Mainer
method:  GET
info:query param

succeed:
	- status code:   200 OK
	- response body: JSON represent get info

failed:
	- status code:   400, 500, ...
	- response body: JSON with error + time
*/

func (h *HttpHandlers) HandleGetInfoMiner(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	statusStr := q.Get("status")
	class := q.Get("class")

	var (
		hasStatus bool
		status    bool
	)

	if statusStr != "" {
		parsed, err := strconv.ParseBool(statusStr)
		if err != nil {
			sendJsonError(w, 400, err.Error())
			return
		}
		hasStatus = true
		status = parsed
	}

	result := h.MinerManager.InfoAll()
	if hasStatus {
		result = h.MinerManager.InfoByStatus(status)
	}

	if class != "" {
		if hasStatus {
			filter := make([]models.MinerStats, 0, len(result))
			for _, miner := range result {
				if miner.Class == class {
					filter = append(filter, miner)
				}
			}

			result = filter
		} else {
			result = h.MinerManager.InfoByGroup(class)
		}
	}
	sendJsonSuccses(w, http.StatusOK, result)

}

/*
pattern: /api/miners/{id}
method:  GET
info:pattern

succeed:
	- status code:   200 OK
	- response body: JSON represent get info

failed:
	- status code:   400, 404, 500, ...
	- response body: JSON with error + time
*/

func (h *HttpHandlers) HandleGetInfoMinerId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	result, err := strconv.Atoi(id)
	if err != nil {
		sendJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	miner, err := h.MinerManager.InfoById(result)
	if err != nil {
		sendJsonError(w, http.StatusNotFound, err.Error())
		return
	}
	sendJsonSuccses(w, http.StatusOK, miner)
}

// /*
// pattern: /api/upgrades
// method:  GET
// info:-

// succeed:
//   - status code:   200 OK
//   - response body: JSON represent get info

// failed:
//   - status code:   400, 500, ...
//   - response body: JSON with error + time
// */

// func (h *HttpHandlers) HandleGetInfoUpdrades(w http.ResponseWriter, r *http.Request) {

// }

/*
pattern: /api/upgrades
method:  Post
info:json

succeed:
	- status code:   200 OK
	- response body: JSON represent get info

failed:
	- status code:   400, 409, 500, ...
	- response body: JSON with error + time
*/

func (h *HttpHandlers) HandleAddUpgrades(w http.ResponseWriter, r *http.Request) {
	var upgradeDTO UpgradeDTO
	if err := json.NewDecoder(r.Body).Decode(&upgradeDTO); err != nil {
		sendJsonError(w, http.StatusBadRequest, err.Error())
		return
	}
	upgrade, err := upgradeDTO.ValidateToUpgradeDTO()
	if err != nil {
		sendJsonError(w, http.StatusBadRequest, err.Error())
	}
	err = h.User.Upgrade(upgrade)
	if err != nil {
		sendJsonError(w, http.StatusBadRequest, err.Error())
		return
	}
	sendJsonSuccses(w, http.StatusAccepted, h.User.GetUpgradeStats())

}

/*
pattern: /api/upgrades?statusBuy=true
method:  GET
info:query param

succeed:
	- status code:   200 OK
	- response body: JSON represent get info

failed:
	- status code:   400, 500, ...
	- response body: JSON with error + time
*/

func (h *HttpHandlers) HandleGetInfoQueryUpdrades(w http.ResponseWriter, r *http.Request) {
	statusParam := r.URL.Query().Get("statusBuy")

	var filterBought *bool
	if statusParam != "" {
		value, err := strconv.ParseBool(statusParam)
		if err != nil {
			sendJsonError(w, http.StatusInternalServerError, err.Error())
			return
		}
		filterBought = &value
	}
	stats := h.User.GetUpgradeStats()
	result := make([]UpgradeDTO, 0, len(stats))

	for name, count := range stats {
		dto := UpgradeDTO{
			Name:      name,
			Count:     count,
			WasBought: count > 0,
		}
		if filterBought != nil && dto.WasBought != *filterBought {
			continue
		}
		result = append(result, dto)
	}

	sendJsonSuccses(w, http.StatusOK, result)
}
