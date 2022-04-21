package integer

import (
	"cbc-casper-go/casper/simulation"
	"fmt"
	"testing"
)

func TestIntegerProtocol_SetInitMsg(t *testing.T) {
	jsonString := simulation.GenerateIntegerJsonString([]float64{1, 2, 3, 4, 5}, []interface{}{100, 50, 1, 0, 24})
	parsedJson, err := parseJson(jsonString)
	if err != nil {
		t.Errorf("error")
		return
	}
	if parsedJson.Conf.Estimates.([]interface{})[0].(float64) != 100 {
		fmt.Println(parsedJson.Conf.Estimates.([]interface{})[0])
		t.Errorf("error")
	}
}
