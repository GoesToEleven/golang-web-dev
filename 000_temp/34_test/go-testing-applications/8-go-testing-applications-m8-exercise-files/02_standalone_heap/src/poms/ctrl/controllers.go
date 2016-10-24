package ctrl

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	home          *HomeController          = new(HomeController)
	shipping      *ShippingController      = new(ShippingController)
	currency      *CurrencyController      = new(CurrencyController)
	purchaseOrder *PurchaseOrderController = new(PurchaseOrderController)
)

func Setup() {
	createResourceServer()

	r := mux.NewRouter()

	r.HandleFunc("/", home.GetHome)
	r.HandleFunc("/api/receivers", shipping.GetReceivers)
	r.HandleFunc("/api/vendors", shipping.GetVendors)
	r.HandleFunc("/api/currencies", currency.GetCurrencies)
	r.HandleFunc("/api/purchaseOrders/{poNumber:\\d+}",
		purchaseOrder.GetOrder).Methods("GET")
	r.HandleFunc("/api/purchaseOrders",
		purchaseOrder.PostOrder).Methods("POST")

	http.Handle("/", r)
}

func createResourceServer() {
	http.Handle("/res/", http.StripPrefix("/res", http.FileServer(http.Dir("res"))))
}
