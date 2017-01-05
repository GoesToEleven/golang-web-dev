package model

import (
	"errors"
	"time"
)

var orders []*PurchaseOrder
var currentOrderId int = 0
var currentDetailId int = 0

type PurchaseOrder struct {
	ID       int         `json:"id"`
	Receiver *Receiver   `json:"receiver"`
	Vendor   *Vendor     `json:"vendor"`
	Currency *Currency   `json:"currency"`
	Purpose  string      `json:"purpose"`
	DueDate  time.Time   `json:"dueDate"`
	Status   string      `json:"status"`
	Details  []*PODetail `json:"details"`
}

type PODetail struct {
	ID            int       `json:"id"`
	Line          int       `json:"line"`
	PartNumber    string    `json:"partNumber"`
	Description   string    `json:"description"`
	Quantity      float32   `json:"quantity"`
	DueDate       time.Time `json:"dueDate"`
	UnitOfMeasure string    `json:"unitOfMeasure"`
	PricePer      float64   `json:"pricePer"`
	Status        string    `json:"status"`
}

func GetOrder(id int) (*PurchaseOrder, error) {

	for _, o := range orders {
		if o.ID == id {
			return o, nil
		}
	}

	return nil, errors.New("Purchase order not found")
}

func SaveOrder(order *PurchaseOrder) *PurchaseOrder {
	order.ID = getNextOrderId()
	order.Status = "NEW"

	for idx := range order.Details {
		order.Details[idx].Line = idx * 10
		order.Details[idx].Status = "NEW"
	}

	orders = append(orders, order)

	return order
}

func getNextOrderId() int {
	currentOrderId++
	return currentOrderId
}

func getNextDetailId() int {
	currentDetailId++
	return currentDetailId
}

func init() {
	receivers := GetReceivers()
	vendors, _ := GetVendors()
	var receiver *Receiver
	if len(receivers) >= 1 {
		receiver = receivers[1]
	} else {
		receiver = &Receiver{}
	}
	var vendor *Vendor
	if len(vendors) >= 2 {
		vendor = vendors[2]
	} else {
		vendor = &Vendor{}
	}

	orders = []*PurchaseOrder{
		{
			ID:       getNextOrderId(),
			Receiver: receiver,
			Vendor:   vendor,
			Status:   "APPROVED",
			Currency: GetCurrencies()[1],
			Purpose:  "Replace leaking valve",
			DueDate:  time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
			Details: []*PODetail{
				{
					ID:            getNextDetailId(),
					Line:          10,
					PartNumber:    "TGY-5412",
					Description:   "Control Valve",
					Quantity:      1,
					DueDate:       time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
					UnitOfMeasure: "EA",
					PricePer:      427.36,
					Status:        "APPROVED",
				}, {
					ID:            getNextDetailId(),
					Line:          20,
					PartNumber:    "M20-2.5x70",
					Description:   "Bolts",
					Quantity:      25,
					DueDate:       time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
					UnitOfMeasure: "EA",
					PricePer:      4.12,
					Status:        "RECEIVED",
				},
			},
		},
	}
}
