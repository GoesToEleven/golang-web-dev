package ctrl

import (
	"encoding/json"
	"net/http"
	"poms/model"
)

type ShippingController struct{}

func (sc *ShippingController) GetReceivers(w http.ResponseWriter, r *http.Request) {
	receivers := model.GetReceivers()

	w.Header().Add("Content-Type", "application/json")

	data, _ := json.Marshal(receivers)

	w.Write(data)
}

func (sc *ShippingController) GetVendors(w http.ResponseWriter, r *http.Request) {
	vendors := model.GetVendors()

	w.Header().Add("Content-Type", "application/json")

	data, _ := json.Marshal(vendors)

	w.Write(data)
}
