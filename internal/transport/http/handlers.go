package http

import (
	"net/http"
	"nilchan/FinalProject/internal/models"
	datamainer "nilchan/FinalProject/internal/service"
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

}

/*
pattern: /api/user/stop
method:  POST
info:    JSON in HTTP request body

succeed:
  - status code:   200 OK
  - response body: JSON represent change status game

failed:
  - status code:   400, 500, ...
  - response body: JSON with error + time
*/

func (h *HttpHandlers) HandleStopGame(w http.ResponseWriter, r *http.Request) {

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
