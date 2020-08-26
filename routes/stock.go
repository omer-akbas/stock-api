package routes

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/omer-akbas/stock-api/models"
	"github.com/omer-akbas/stock-api/utils"
)

func stockDetailsGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	if code == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	data, err := models.StockByCode(strings.ToUpper(code))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, data)
}

func stockRatesGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	if code == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	data, err := models.StockRatesByCode(strings.ToUpper(code))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"status": data})
}

func stockListRatesGet(w http.ResponseWriter, r *http.Request) {
	data, err := models.StockListRates()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, data)
}
