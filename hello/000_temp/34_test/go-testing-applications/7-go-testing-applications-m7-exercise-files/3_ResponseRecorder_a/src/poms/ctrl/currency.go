package ctrl

import (
	"encoding/json"
	"net/http"
	"poms/model"
)

type CurrencyController struct{}

func (cc *CurrencyController) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies := model.GetCurrencies()

	data, _ := json.Marshal(currencies)

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
