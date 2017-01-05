package model

type Currency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetCurrencies() []*Currency {
	result := []*Currency{
		{ID: 1, Name: "USD"},
		{ID: 2, Name: "EUR"},
	}

	return result
}
