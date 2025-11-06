package http

import (
	"encoding/json"
	"net/http"
	"nilchan/FinalProject/internal/models"
	datamainer "nilchan/FinalProject/internal/service"
	"nilchan/FinalProject/internal/service/dataminer/minimainer"
	"nilchan/FinalProject/internal/service/dataminer/normalmainer"
	"nilchan/FinalProject/internal/service/dataminer/strongminer"
	"time"
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
		http.Error(w, err.Error(), code)
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
	message := h.MinerManager.StopGame()
	sendJsonSuccses(w, http.StatusOK, message)

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

}
