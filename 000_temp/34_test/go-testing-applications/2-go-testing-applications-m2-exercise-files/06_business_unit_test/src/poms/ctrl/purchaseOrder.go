package ctrl

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"poms/model"
	"strconv"
)

type PurchaseOrderController struct{}

func (poc *PurchaseOrderController) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	poNumber, _ := strconv.Atoi(vars["poNumber"])

	order, err := model.GetOrder(poNumber)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	data, _ := json.Marshal(order)
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (poc *PurchaseOrderController) PostOrder(w http.ResponseWriter, r *http.Request) {
	var order *model.PurchaseOrder
	data := make([]byte, r.ContentLength)
	r.Body.Read(data)

	err := json.Unmarshal(data, &order)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	order = model.SaveOrder(order)

	resp, _ := json.Marshal(order)
	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}
