package routes

import (
	"net/http"

	"github.com/omer-akbas/stock-api/utils"
)

func indexGet(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"result": "stock api"})
}
