package main

import (
	"testing"
)

func TestTempConvert(t *testing.T) {
	get1, get2 := convertKelvin(100)
	var res1, res2 float64
	res1, res2 = -279.66999999999996, -173.14999999999998
	if get1 != res1 {
		t.Errorf("Got: %f Wanted: %f", res1, get1)
	}
	if get2 != res2 {
		t.Errorf("Got: %f Wanted: %f", res2, get2)
	}
	get1, get2 = convertCelsius(100)
	res1, res2 = 212, 373.15
	if get1 != res1 {
		t.Errorf("Got: %f Wanted: %f", res1, get1)
	}
	if get2 != res2 {
		t.Errorf("Got: %f Wanted: %f", res2, get2)
	}
	get1, get2 = convertFahrenheit(100)
	res1, res2 = 310.92777777777775, 37.77777777777778
	if get1 != res1 {
		t.Errorf("Got: %f Wanted: %f", res1, get1)
	}
	if get2 != res2 {
		t.Errorf("Got: %f Wanted: %f", res2, get2)
	}
}
