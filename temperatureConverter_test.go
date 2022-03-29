package main

import (
	"testing"
)

func TestTempConvert(t *testing.T) {
	get1, get2 := convertKelvin(100)
	var res1, res2 float64
	res1, res2 = -279.67, -173.150000
	if get1 != res1 {
		t.Errorf("Got: %f Wanted: %f", res1, get1)
	}
	if get2 != res2 {
		t.Errorf("Got: %f Wanted: %f", res2, get2)
	}
	get1, get2 = convertCelsius(100)
	res1, res2 = 212, 373.15000000000003
	if get1 != res1 {
		t.Errorf("Got: %f Wanted: %f", res1, get1)
	}
	if get2 != res2 {
		t.Errorf("Got: %f Wanted: %f", res2, get2)
	}
	get1, get2 = convertFahrenheit(100)
	res1, res2 = 310.9277777777778, 37.77777777777778
	if get1 != res1 {
		t.Errorf("Got: %f Wanted: %f", res1, get1)
	}
	if get2 != res2 {
		t.Errorf("Got: %f Wanted: %f", res2, get2)
	}
}
