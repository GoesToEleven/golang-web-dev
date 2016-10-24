package model

import (
	"encoding/json"
	"testing"
)

func TestGetsTheCorrectCurencies(t *testing.T) {
	//arrange

	//act
	result := GetCurrencies()

	//assert
	resultJSON, _ := json.Marshal(result)
	expected, _ := json.Marshal([]*Currency{
		{ID: 1, Name: "USD"},
		{ID: 2, Name: "EUR"},
	})

	if string(resultJSON) != string(expected) {
		t.Error("Did not get expected currencies")
	}

}
